package services

import (
	"PPO_BMSTU/internal/services/service_interfaces"
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
)

type TwoFAService struct {
	client *api.Client
}

// UserCode представляет структуру для хранения кода 2FA.
type UserCode struct {
	Code       string    `json:"code"`
	Expiration time.Time `json:"expiration"`
}

// NewTwoFAService создает новый сервис для работы с Vault.
func NewTwoFAService(vaultAddr, token string) service_interfaces.ITwoFA {
	config := api.DefaultConfig()
	config.Address = vaultAddr

	client, err := api.NewClient(config)
	if err != nil {
		return nil
	}

	client.SetToken(token)
	return &TwoFAService{client: client}
}

// GenerateAndStoreCode генерирует 2FA-код и сохраняет его в Vault.
func (s *TwoFAService) GenerateAndStoreCode(userID string) (string, error) {
	code := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000) // Генерация 6-значного кода
	expiration := time.Now().Add(5 * time.Minute)              // Код действует 5 минут

	secretData := map[string]interface{}{
		"code":       code,
		"expiration": expiration.Format(time.RFC3339),
	}

	// Сохраняем код в Vault
	_, err := s.client.Logical().Write(fmt.Sprintf("secret/data/2fa/%s", userID), map[string]interface{}{
		"data": secretData,
	})
	if err != nil {
		return "", err
	}

	return code, nil
}

// GetCode получает код 2FA для указанного пользователя из Vault.
func (s *TwoFAService) GetCode(userID string) (*UserCode, error) {
	secret, err := s.client.Logical().Read(fmt.Sprintf("secret/data/2fa/%s", userID))
	if err != nil {
		return nil, err
	}

	if secret == nil || secret.Data["data"] == nil {
		return nil, fmt.Errorf("code not found for user %s", userID)
	}

	data := secret.Data["data"].(map[string]interface{})
	code := data["code"].(string)
	expiration, _ := time.Parse(time.RFC3339, data["expiration"].(string))

	return &UserCode{Code: code, Expiration: expiration}, nil
}

// VerifyCode проверяет введенный пользователем код.
func (s *TwoFAService) VerifyCode(userID, code string) (bool, error) {
	userCode, err := s.GetCode(userID)
	if err != nil {
		return false, err
	}

	// Проверяем, что код совпадает и не истек срок действия
	if userCode.Code == code && userCode.Expiration.After(time.Now()) {
		return true, nil
	}

	return false, nil
}

//
//func main() {
//	// Настройки Vault
//	vaultAddr := "http://localhost:8200"
//	token := "root"
//
//	// Создаем сервис 2FA
//	service, err := NewTwoFAService(vaultAddr, token)
//	if err != nil {
//		panic(err)
//	}
//
//	// Пример работы
//	userID := "user123"
//	code, err := service.GenerateAndStoreCode(userID)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Generated 2FA code for user %s: %s\n", userID, code)
//
//	// Проверяем код
//	isValid, err := service.VerifyCode(userID, code)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Is code valid? %v\n", isValid)
//}

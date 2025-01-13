package services

import (
	"PPO_BMSTU/internal/services/service_interfaces"
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
)

// UserCode представляет структуру для хранения кода 2FA.
type UserCode struct {
	Code       string    `json:"code"`
	Expiration time.Time `json:"expiration"`
}

type TwoFAService struct {
	client      *api.Client
	emailConfig EmailConfig // Конфигурация для отправки email
}

type EmailConfig struct {
	SMTPHost    string
	SMTPPort    int
	Username    string
	Password    string
	FromAddress string
}

// NewTwoFAService создает новый сервис для работы с Vault и отправкой email.
func NewTwoFAService(vaultAddr, token string) service_interfaces.ITwoFA {
	var emailConfig = EmailConfig{
		SMTPHost:    "smtp.example.com",
		SMTPPort:    587,
		Username:    "your-email@example.com",
		Password:    "your-password",
		FromAddress: "no-reply@example.com",
	}

	config := api.DefaultConfig()
	config.Address = vaultAddr

	client, err := api.NewClient(config)
	if err != nil {
		return nil
	}

	client.SetToken(token)
	return &TwoFAService{client: client, emailConfig: emailConfig}
}

// GenerateAndStoreCode генерирует 2FA-код, сохраняет его в Vault и отправляет на email.
func (s *TwoFAService) GenerateAndStoreCode(userID, email string) (string, error) {
	// Генерация 6-значного кода
	code := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
	expiration := time.Now().Add(5 * time.Minute) // Код действует 5 минут

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

	// Отправка кода на email
	//err = s.sendEmail(email, code)
	//if err != nil {
	//	return "", fmt.Errorf("failed to send email: %w", err)
	//}

	return code, nil
}

// sendEmail отправляет 2FA-код на указанный email.
//func (s *TwoFAService) sendEmail(email, code string) error {
//	auth := smtp.PlainAuth("", s.emailConfig.Username, s.emailConfig.Password, s.emailConfig.SMTPHost)
//	to := []string{email}
//	subject := "Your 2FA Code"
//	body := fmt.Sprintf("Your 2FA code is: %s. It will expire in 5 minutes.", code)
//
//	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", email, subject, body))
//
//	address := fmt.Sprintf("%s:%d", s.emailConfig.SMTPHost, s.emailConfig.SMTPPort)
//	return smtp.SendMail(address, auth, s.emailConfig.FromAddress, to, message)
//}

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

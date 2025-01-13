package tests

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/server/api/modelsViewApi"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (suite *e2eTestSuite) aJudgeExistsWithEmailAndPassword(email, password string) (string, error) {
	// Используем структуру JudgeInput для создания нового пользователя
	user := modelsViewApi.JudgeInput{
		FIO:      "Test User2",
		Login:    email,
		Password: password,
		Role:     1,
		Post:     "Judge",
	}

	reqBody, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user data: %s", err)
	}

	req, _ := http.NewRequest("POST", "/api/judges/", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	// Проверка успешного создания пользователя
	if resp.Code != http.StatusCreated {
		return "", fmt.Errorf("failed to create user: %s", resp.Body.String())
	}

	// Структура для парсинга ответа
	var response models.Judge

	// Парсим ответ в структуру
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %s", err)
	}

	// Возвращаем ID пользователя
	return response.ID.String(), nil
}

func (suite *e2eTestSuite) theUserLogsInWithEmailAndPassword(email, password string) error {
	// Используем структуру LoginFormData для логина
	loginData := modelsViewApi.LoginFormData{
		Login:    email,
		Password: password,
	}

	reqBody, err := json.Marshal(loginData)
	if err != nil {
		return fmt.Errorf("failed to marshal login data: %s", err)
	}

	req, _ := http.NewRequest("POST", "/api/signin", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		return fmt.Errorf("failed to log in: %s", resp.Body.String())
	}

	return nil
}

func (suite *e2eTestSuite) theUserGeneratesA2FACode(judgeID string) (string, error) {
	url := fmt.Sprintf("/api/2fa/%s/generate", judgeID)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Accept", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		return "", fmt.Errorf("failed to generate 2FA code: %s", resp.Body.String())
	}

	var generateResp struct {
		Code string `json:"code"`
	}
	err := json.Unmarshal(resp.Body.Bytes(), &generateResp)
	if err != nil {
		return "", fmt.Errorf("failed to parse 2FA code: %s", err)
	}
	return generateResp.Code, nil
}

func (suite *e2eTestSuite) theUserVerifiesThe2FACode(code, judgeID string) error {
	// Отправляем код в запросе
	verifyReqBody := fmt.Sprintf(`{
		"code": "%s"
	}`, code)
	url := fmt.Sprintf("/api/2fa/%s/verify", judgeID)

	req, _ := http.NewRequest("POST", url, strings.NewReader(verifyReqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		return fmt.Errorf("failed to verify 2FA code: %s", resp.Body.String())
	}
	return nil
}

func (suite *e2eTestSuite) theUserChangesTheirPasswordTo(newPassword, email string) error {
	// Отправляем новый пароль
	updatePasswordReqBody := fmt.Sprintf(`{
		"password": "%s"
	}`, newPassword)
	url := fmt.Sprintf("/api/judges/%s/password", email)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(updatePasswordReqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		return fmt.Errorf("failed to change password: %s", resp.Body.String())
	}
	return nil
}

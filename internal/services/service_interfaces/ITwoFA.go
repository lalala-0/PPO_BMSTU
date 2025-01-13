package service_interfaces

// TwoFAServiceInterface определяет контракт для сервиса 2FA.
type ITwoFA interface {
	GenerateAndStoreCode(userID string) (string, error)
	VerifyCode(userID, code string) (bool, error)
}

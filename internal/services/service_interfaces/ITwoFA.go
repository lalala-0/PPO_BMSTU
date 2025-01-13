package service_interfaces

// ITwoFA определяет контракт для сервиса 2FA.
type ITwoFA interface {
	// GenerateAndStoreCode генерирует и сохраняет 2FA-код, а также отправляет его на указанный email.
	GenerateAndStoreCode(userID, email string) (string, error)

	// VerifyCode проверяет введённый пользователем 2FA-код.
	VerifyCode(userID, code string) (bool, error)
}

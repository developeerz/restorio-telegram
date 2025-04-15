package telegram

type Cache interface {
	PutUser(telegram string, userJSON []byte) error
	PutVerificationCode(telegram string, code int) error
	GetUser(telegram string) ([]byte, error)
	GetVerificationCode(telegram string) (int, error)
}

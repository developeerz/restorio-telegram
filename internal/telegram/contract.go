package telegram

type Repository interface {
	GetCodeByTelegram(telegram string) (int, error)
	UpdateUserByTelegram(telegram string, telegramID int64) error
}

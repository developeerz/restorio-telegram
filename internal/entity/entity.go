package entity

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Telegram   string `gorm:"unique"`
	TelegramID int64
}

type UserCode struct {
	Telegram string `gorm:"unique"`
	Code     int
}

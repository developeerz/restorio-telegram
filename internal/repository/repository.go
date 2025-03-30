package repository

import (
	"github.com/developeerz/restorio-telegram/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetCodeByTelegram(telegram string) (int, error)
	UpdateUserByTelegram(telegram string, telegramID int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRipository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetCodeByTelegram(telegram string) (int, error) {
	var userCode models.UserCode
	result := r.db.Where("telegram = ?", telegram).First(&userCode)
	return userCode.Code, result.Error
}

func (r *userRepository) UpdateUserByTelegram(telegram string, telegramID int64) error {
	return r.db.Table("users").Where("telegram = ?", telegram).Update("telegram_id", telegramID).Error
}

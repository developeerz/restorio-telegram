package repository

import (
	"github.com/developeerz/restorio-telegram/internal/repository/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetCodeByTelegram(telegram string) (int, error) {
	var userCode models.UserCode
	result := r.db.Where("telegram = ?", telegram).First(&userCode)

	return userCode.Code, result.Error
}

func (r *Repository) UpdateUserByTelegram(telegram string, telegramID int64) error {
	return r.db.Table("users").Where("telegram = ?", telegram).Update("telegram_id", telegramID).Error
}

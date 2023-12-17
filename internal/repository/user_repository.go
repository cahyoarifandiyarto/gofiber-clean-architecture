package repository

import (
	"github.com/cahyoarifandiyarto/gofiber-clean-architecture/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) CountByUsername(db *gorm.DB, username string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("username = ?", username).Count(&total).Error
	return total, err
}

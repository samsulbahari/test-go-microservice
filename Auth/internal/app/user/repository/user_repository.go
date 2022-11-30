package repository

import (
	"auth/internal/app/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}

func (ur UserRepository) CreateUser(ctx context.Context, user domain.User) error {
	err := ur.db.WithContext(ctx).Create(&user).Error
	return err
}

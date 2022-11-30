package service

import (
	"auth/internal/app/domain"
	"auth/internal/app/libraries"
	"context"
	"errors"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo}
}

func (us UserService) Login(ctx context.Context, param domain.Login) (string, error) {
	user, err := us.userRepo.GetByUsername(ctx, param.Username)
	if err != nil {
		return "", errors.New("wrong email/password")
	}

	err = user.ComparePassword(param.Password)
	if err != nil {
		return "", errors.New("wrong email/password")
	}

	return libraries.GenerateJWT(user.ID)
}

func (us UserService) Register(ctx context.Context, user domain.User) (int, error) {

	username, _ := us.userRepo.GetByUsername(ctx, user.Username)
	if username.Username != "" {
		return 200, errors.New("Username Already exists")
	}
	user.Role = "user"
	user.Password = user.HashPassword(user.Password)
	err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		return 400, errors.New("Server error")
	}

	return 200, nil

}

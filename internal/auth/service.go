package auth

import (
	"errors"
	"go-advance/internal/user"
)

type AuthService struct {
	repository *user.UserRepository
}

func NewUserService(repository *user.UserRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	exitstedUser, _ := service.repository.FindByEmail(email)
	if exitstedUser != nil {
		return "", errors.New(ErrUserAlreadyExist)
	}

	user := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}
	err := service.repository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

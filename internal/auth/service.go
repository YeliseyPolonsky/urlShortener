package auth

import (
	"errors"
	"go-advance/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository *user.UserRepository
}

func NewUserService(repository *user.UserRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (service *AuthService) Login(email, password string) error {
	exitstedUser, _ := service.repository.FindByEmail(email)

	if exitstedUser == nil {
		return errors.New(ErrUserNotRegistered)
	}

	err := bcrypt.CompareHashAndPassword([]byte(exitstedUser.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	exitstedUser, _ := service.repository.FindByEmail(email)
	if exitstedUser != nil {
		return "", errors.New(ErrUserAlreadyExist)
	}

	cryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Email:    email,
		Password: string(cryptedPass),
		Name:     name,
	}
	err = service.repository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

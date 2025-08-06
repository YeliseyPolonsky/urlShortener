package auth_test

import (
	"go-advance/internal/auth"
	"go-advance/internal/user"
	"testing"
)

type MockUserRepositoty struct{}

func (r *MockUserRepositoty) Create(user *user.User) error {
	return nil
}

func (r *MockUserRepositoty) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}
func TestSuccessRegister(t *testing.T) {
	err := auth.NewUserService(&MockUserRepositoty{}).Register(
		"example@gmail.com",
		"123",
		"Bacя",
	)

	if err != nil {
		t.Fatal(err)
	}
}

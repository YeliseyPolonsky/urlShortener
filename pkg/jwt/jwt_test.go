package jwt_test

import (
	"go-advance/pkg/jwt"
	"testing"
)

func TestJWTService(t *testing.T) {
	const email = "example@gmail.com"
	jwtService := jwt.NewJWT("/2+XnmJGz1j3ehIVI/5P9kl+CghrE3DcS7rnT+qar5w=")
	token, err := jwtService.Create(email)
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatalf("Token is invalid %v", token)
	}
	if data.Email != email {
		t.Fatalf("Expected %v got %v", email, data.Email)
	}
}

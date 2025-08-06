package auth_test

import (
	"bytes"
	"encoding/json"
	"go-advance/configs"
	"go-advance/internal/auth"
	"go-advance/internal/user"
	"go-advance/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: database}))
	if err != nil {
		return nil, nil, err
	}

	userRepository := &user.UserRepository{Db: &db.Db{DB: gormDB}}
	authService := auth.NewAuthService(userRepository)
	handler := &auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "/2+XnmJGz1j3ehIVI/5P9kl+CghrE3DcS7rnT+qar5w=",
			},
		},
		AuthService: authService,
	}

	return handler, mock, nil
}

func TestSuccessLogin(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow(
			"yelisey@gmail.com",
			"$2a$10$uAhFvyQXskxyoAArPac2XeocCQQHAGMPz/t9YGOGwDH6TpRlaxLtG",
		)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "yelisey@gmail.com",
		Password: "123",
	})
	reader := bytes.NewReader(data)
	writer := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(writer, req)
	if writer.Code != http.StatusOK {
		t.Fatal(writer.Code)
	}
}

func TestSuccessReg(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}
	table := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(table)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	data, _ := json.Marshal(auth.RegisterRequest{
		Name:     "yelisey",
		Email:    "yelisey@gmail.com",
		Password: "123",
	})
	reader := bytes.NewReader(data)
	writer := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(writer, req)
	if writer.Code != http.StatusOK {
		t.Fatal(writer.Code)
	}
}

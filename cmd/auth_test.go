package main

import (
	"bytes"
	"encoding/json"
	"go-advance/internal/auth"
	"go-advance/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func InitData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "yelisey@gmail.com",
		Password: "$2a$10$uAhFvyQXskxyoAArPac2XeocCQQHAGMPz/t9YGOGwDH6TpRlaxLtG",
		Name:     "Василий",
	})
}

func RemoveData(db *gorm.DB) {
	db.Unscoped().Where("email = ?", "yelisey@gmail.com").Delete(&user.User{})
}

func TestSuccessLogin(t *testing.T) {
	//Prepare
	db := InitDb()
	InitData(db)

	s := httptest.NewServer(App())

	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "yelisey@gmail.com",
		Password: "123",
	})

	res, err := http.Post(s.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var body auth.LoginResponse
	err = json.Unmarshal(data, &body)
	if err != nil {
		t.Fatal(err)
	}
	if body.Token == "" {
		t.Fatal("TOKEN IS EMPTY ", err)
	}

	RemoveData(db)
}

func TestFailLogin(t *testing.T) {
	//Prepare
	db := InitDb()
	InitData(db)

	s := httptest.NewServer(App())

	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "yelisey@gmail.com",
		Password: "1",
	})

	res, err := http.Post(s.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode == 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}

	RemoveData(db)
}

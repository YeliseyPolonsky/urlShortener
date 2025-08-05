package req

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HandleBody[T any](r *http.Request) (*T, error) {
	dto, err := Decode[T](r.Body)
	if err != nil {
		return nil, err
	}
	err = Validate(dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func Decode[T any](body io.ReadCloser) (*T, error) {
	var dto T
	err := json.NewDecoder(body).Decode(&dto)
	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func Validate[T any](dto *T) error {
	validator := validator.New()
	err := validator.Struct(dto)
	return err
}

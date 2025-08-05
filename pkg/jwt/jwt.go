package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	secret string
}

type JWTData struct {
	Email string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (j *JWT) Create(email string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return false, nil
	}
	data := &JWTData{
		Email: t.Claims.(jwt.MapClaims)["email"].(string),
	}
	return t.Valid, data
}

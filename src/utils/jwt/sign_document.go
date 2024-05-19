package jwt

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/config"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Id uuid.UUID `json:"id"`
}

func (c *CustomClaims) Valid() error {
	return nil
}

func NewCustomClaims(id uuid.UUID) *CustomClaims {
	return &CustomClaims{
		Id: id,
	}
}

func SignDocument(id uuid.UUID) string {
	claims := NewCustomClaims(id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	envs := config.LoadEnvs(".env")
	secret := []byte(envs.Get("JWT_SECRET"))
	signedToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	return signedToken
}

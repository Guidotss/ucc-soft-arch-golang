package jwt

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/config"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Id   uuid.UUID `json:"id"`
	Role int       `json:"role"`
}

func (c *CustomClaims) Valid() error {
	return nil
}

func NewCustomClaims(id uuid.UUID, role int) *CustomClaims {
	return &CustomClaims{
		Id:   id,
		Role: role,
	}
}

func SignDocument(id uuid.UUID, role int) string {
	claims := NewCustomClaims(id, role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	envs := config.LoadEnvs(".env")
	secret := []byte(envs.Get("JWT_SECRET"))
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return ""
	}
	return signedToken
}

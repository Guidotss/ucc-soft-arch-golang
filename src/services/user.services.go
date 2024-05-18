package services

import "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/auth"

func LoginService(auth.LoginRequest) *auth.LoginResponse {
	return &auth.LoginResponse{
		Username: "Guido",
		Message:  "Login",
		Token:    "123456789",
	}
}

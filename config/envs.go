package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs interface {
	Get(key string) string
}

type envsImpl struct{}

func (envsImpl) Get(key string) string {
	return os.Getenv(key)
}

func LoadEnvs(filename ...string) Envs {
	err := godotenv.Load(filename...)
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	return &envsImpl{}
}

package config

import (
	"fmt"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(model.User{}, model.Course{}, model.Categories{}, model.Inscriptos{})

	return db
	// defer db.Close()
}

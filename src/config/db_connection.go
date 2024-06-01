package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	//db.AutoMigrate(model.User{}, model.Course{}, model.Categories{}, model.Inscripto{}, model.Ratings{}, model.Comments{})

	return db
	// defer db.Close()
}

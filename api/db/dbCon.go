package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func Connect() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	if db == nil {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print("Connected")
		}
	}
	return db
}

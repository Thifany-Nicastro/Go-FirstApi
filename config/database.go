package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connection created successfully!")

	return db
}

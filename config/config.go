package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

// var (
// 	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_PORT"),
// 		os.Getenv("DB_NAME"))
// )

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
}

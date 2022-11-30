package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host= %s user= %s password= %s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", db_host, db_username, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error Connect database")
	}
	return db
}

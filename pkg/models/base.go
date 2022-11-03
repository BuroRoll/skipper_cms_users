package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var cmsDb *gorm.DB

func init() {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable port=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_PASSWORD"))
	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatalf("error %s", err)
	}
	cmsDb = conn

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func GetDB() *gorm.DB {
	return cmsDb
}

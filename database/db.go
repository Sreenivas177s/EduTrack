package database

import (
	"chat-server/api-framework/entity"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDBRef() *gorm.DB {
	if dbInstance == nil {
		InitDataBase()
	}
	return dbInstance
}

func InitDataBase() {
	dbParams := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_TIMEZONE"))
	gormConfig := &gorm.Config{}
	dbref, err := gorm.Open(postgres.Open(dbParams), gormConfig)
	if err != nil {
		panic("Error while obtaining db instance")
	}
	dbInstance = dbref

	// migrate tables
	if os.Getenv("GORM_MIGRATE") == "true" {
		handleTableMigration(dbInstance)
	}
}

func handleTableMigration(db *gorm.DB) {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err.Error())
	}
}

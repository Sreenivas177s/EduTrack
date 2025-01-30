package database

import (
	"chat-server/api/entity"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDBRef() *gorm.DB {
	if dbInstance == nil {
		initDataBase()
	}
	return dbInstance
}

func Init() {
	initDataBase()
	initializeRedis()
}

func initDataBase() {
	dbParams := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_TIMEZONE"))
	gormConfig := &gorm.Config{}
	dbref, err := gorm.Open(postgres.Open(dbParams), gormConfig)
	if err != nil {
		panic("Error while obtaining db instance")
	}
	log.Info("Connected to Database")
	dbInstance = dbref

	// migrate tables
	if os.Getenv("GORM_MIGRATE") == "true" {
		handleTableMigration(dbInstance)
	}
}

func handleTableMigration(db *gorm.DB) {
	//user tables
	MigrateStruct(&entity.User{})
	// panicOnError(db.-(&entity.Institution{}))
	// panicOnError(db.AutoMigrate(&entity.Campus{}))

}
func MigrateStruct(structs ...interface{}) {
	db := GetDBRef()
	if err := db.AutoMigrate(structs...); err != nil {
		panic(err.Error())
	}
}

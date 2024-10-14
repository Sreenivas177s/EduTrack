package database

import (
	"github.com/jmoiron/sqlx"
)

var dbInstance *sqlx.DB

func GetDBRef() any {
	if dbInstance == nil {
		InitDataBase()
	}
	return dbInstance
}

func InitDataBase() {
	// connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	// dbInstance = sqlx.MustConnect("pgx", connectionString)
}

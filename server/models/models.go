package models

import (
	"github.com/tboerc/gwall/shared"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	models []interface{}
)

func Connect() (err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: shared.Getenv("DATABASE_URL", ""),
	}))
	if err != nil {
		return
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(models...)

	return
}

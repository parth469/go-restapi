package utils

import (
	"errors"

	"github.com/parth469/go-restapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB(dbs string) error {
	db, err := gorm.Open(postgres.Open(dbs), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return errors.New(err.Error())
	}

	Database = DBInstance{DB: db}

	db.AutoMigrate(&models.User{})

	return nil
}

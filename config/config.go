package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnviroment() {
	if err := godotenv.Load("app.env"); err != nil {
		logrus.Fatalf("Error loading .env file : %s", err)
	}
}

func SetUpDatabase() (*gorm.DB, error) {
	dsn := os.Getenv(DatabaseUrl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connection to db: %s", err)
	}

	return db, nil

}

func RunAutoMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(
	//fill models
	); err != nil {
		errorMessage := fmt.Sprintf("Error migrating database %s", err)
		logrus.Error(errorMessage)
		return errors.New(errorMessage)
	}

	return nil

}

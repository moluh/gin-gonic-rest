package config

import (
	"os"

	model "github.com/moluh/ginrest/model"
	util "github.com/moluh/ginrest/util"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("ENV") != "production" {
		databaseURI <- util.GodotEnv("SQL_DATABASE_DEV")
	} else {
		databaseURI <- os.Getenv("SQL_DATABASE_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.UserModel{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}

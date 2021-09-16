package migrations

import (
	"fmt"

	"../models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const dbinfo = "host=localhost user=postgres password=12345 dbname=BigDb port=5432 sslmode=disable"

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	DB.AutoMigrate(&models.User{})

}

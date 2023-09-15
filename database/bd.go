package database

import (
	"fmt"
	"giftjob-backend/graph/model"
	"giftjob-backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	host := utils.Getenv("DB_HOST")
	user := utils.Getenv("DB_USER")
	password := utils.Getenv("DB_PASSWORD")
	dbname := utils.Getenv("DB_NAME")
	port := utils.Getenv("DB_PORT")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := DB.AutoMigrate(&model.User{}, &model.CompanyCustomTemplate{}, &model.Company{}, &model.CompanyCustomField{}, &model.Person{}, &model.Note{}, &model.Calendar{}); err != nil {
		panic("failed to migrate database")
	}
}

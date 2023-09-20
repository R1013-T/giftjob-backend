package database

import (
	"giftjob-backend/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func Init() {
	url := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := DB.AutoMigrate(&model.User{}, &model.CompanyCustomTemplate{}, &model.CompanyCustomTemplateField{}, &model.Company{}, &model.CompanyCustomField{}, &model.Person{}, &model.Note{}, &model.Calendar{}); err != nil {
		panic("failed to migrate database")
	}
}

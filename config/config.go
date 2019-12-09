package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-training-restful/models"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

	var err error
	DB, err = gorm.Open("sqlite3", "./lib/contohFile.db")
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
}

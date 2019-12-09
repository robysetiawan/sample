package database

import (
	"go-training-restful/config"
	"go-training-restful/models"
)

var db = config.DB

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

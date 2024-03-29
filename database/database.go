package database

import (
	"goV2Web/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DataBaseName = "users.db"

// Initialise Database
func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open(DataBaseName), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Users{})
	return nil
}

// This method create new User in Database
func CreateUser(email string, active bool) (models.Users, error) {
	var newUsers = models.Users{Email: email, Active: active}
	db, err := gorm.Open(sqlite.Open(DataBaseName), &gorm.Config{})
	if err != nil {
		return newUsers, err
	}
	db.Create(&newUsers)

	return newUsers, nil
}

// This method returns all users that stored in db
func GetAllUsers() ([]models.Users, error) {
	var users []models.Users

	db, err := gorm.Open(sqlite.Open(DataBaseName), &gorm.Config{})
	if err != nil {
		return users, err
	}

	db.Find(&users)

	return users, nil
}

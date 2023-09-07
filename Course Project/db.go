package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbFile = "chats.db"
	DB     *gorm.DB
)

type User struct {
	gorm.Model // This adds fields: ID, CreatedAt, UpdatedAt, DeletedAt
	TelegramID int
	Username   string
	Language   string

	// or use telegram id as primary key
}

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate database")
	}

	DB = db
	return nil
}

func CreateUser(user User) (*User, error) {
	result := DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindUserByUsername(username string) (*User, error) {
	var user User
	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindUserByID(id int) (*User, error) {
	var user User
	result := DB.Where("telegram_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

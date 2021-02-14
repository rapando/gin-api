package models

import "github.com/rapando/gin-api/config"

// GetAllUsers fetch all user data
func GetAllUsers(user *[]User) error {
	return config.DB.Find(user).Error
}

// CreateUser model
func CreateUser(user *User) error {
	return config.DB.Create(user).Error
}

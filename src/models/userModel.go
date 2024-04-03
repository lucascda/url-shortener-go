package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateUser struct {
	Email                string
	Name                 string
	Password             string
	PasswordConfirmation string
}

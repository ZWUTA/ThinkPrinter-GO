package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;"`
	Password string `gorm:"not null"`
	Name     string
	Number   string
	Vip      bool `gorm:"default:false"`
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Number   string `json:"number"`
}

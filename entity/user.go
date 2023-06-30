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

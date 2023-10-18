package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;"`
	Password string `gorm:"not null"`
}

func (u *User) ToLoginForm() LoginForm {
	return LoginForm{
		Username: u.Username,
		Password: u.Password,
	}
}

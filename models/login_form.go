package models

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (lf *LoginForm) ToUser() User {
	return User{
		Username: lf.Username,
		Password: lf.Password,
	}
}

package models

type User struct {
	Fullname        string `form:"fullname" json:"fullname"`
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword"`
}

var Users []User

type AuthCredentials struct {
		Fullname string `form:"fullname" json:"fullname"`
		Email    string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
	}
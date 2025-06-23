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

var OTP string

type ResetPassword struct {
	Email           string `form:"email" json:"email"`
	Otp             string `form:"otp" json:"otp"`
	NewPassword     string `form:"newPassword" json:"newPassword"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword"`
}

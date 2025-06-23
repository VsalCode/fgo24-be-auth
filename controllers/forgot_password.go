package controllers

import (
	"fmt"
	"gin-authflow/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(ctx *gin.Context) {
	var email string
	ctx.ShouldBind(&email)

	randomNumber := rand.Intn(900000) + 100000

	isEmailValid := "valid"
	for _, item := range models.Users {
		if email == item.Email {
			isEmailValid = "valid"
			break
		}
		isEmailValid = "not valid"
	}

	if isEmailValid == "not valid" {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "your email is not found!",
		})
		return
	}

	models.OTP = randomNumber
	fmt.Println(randomNumber)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "OTP anda telah dikirimkan",
		Result:  models.OTP,
	})

}

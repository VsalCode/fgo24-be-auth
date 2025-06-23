package controllers

import (
	"gin-authflow/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResetPassword(ctx *gin.Context) {
	input := models.ResetPassword{}
	ctx.ShouldBind(&input)

	var updatedUser models.User
	found := false
	for i, item := range models.Users {
		if input.Email == item.Email && input.Otp == models.OTP {
			models.Users[i].Password = input.NewPassword
			updatedUser = models.Users[i]
			found = true
			break
		}
	}

	if !found {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "your email or OTP is wrong!",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Data Successfully Updated!",
		Result:  updatedUser,
	})
}

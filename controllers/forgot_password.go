package controllers

import (
	"gin-authflow/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(ctx *gin.Context) {
	var email string

	ctx.ShouldBind(&email)

	for _, item := range models.Users {
		if email == item.Email {
			break
		}

	}

	ctx.JSON(http.StatusBadRequest, models.Response{
		Success: false,
		Message: "Login Failed! your account is not found!",
	})

}

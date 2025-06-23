package controllers

import (
	"gin-authflow/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthLogin(ctx *gin.Context) {
	credentials := models.AuthCredentials{}

	ctx.ShouldBind(&credentials)

	statusLogin := "login"
	for _, item := range models.Users {
		if credentials.Email == item.Email && credentials.Password == item.Password {
			statusLogin = "login"
			break
		}
		statusLogin = "cannot login"
	}

	if statusLogin == "cannot login" {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Login Failed! your account is not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login success",
		Result:  credentials,
	})

}

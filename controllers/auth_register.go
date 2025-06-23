package controllers;

import (
	"gin-authflow/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func AuthRegister(ctx *gin.Context) {
		tempData := models.User{}

		ctx.ShouldBind(&tempData)

		if tempData.Fullname == "" && tempData.Email == "" && tempData.Password == "" && tempData.ConfirmPassword == "" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Field input can't be empty!",
			})
			return
		}

		if tempData.Password != tempData.ConfirmPassword {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "password must be match with confirm password!",
			})
			return
		}

		status := "available"
		if len(models.Users) > 0 {
			for _, item := range models.Users {
				if tempData.Email == item.Email {
					status = "not available"
					break
				}
				status = "available"
			}

			if status == "not available" {
				ctx.JSON(http.StatusBadRequest, models.Response{
					Success: false,
					Message: "Email is not available!",
				})
				return
			}
		}

		models.Users = append(models.Users, tempData)

		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "register succes",
			Result:  models.Users,
		})

	}
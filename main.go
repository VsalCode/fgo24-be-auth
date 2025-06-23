package main

import (
	"gin-authflow/models"
	"gin-authflow/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.CombineRouter(r)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "Login Success",
		})
	})

	r.Run()
}

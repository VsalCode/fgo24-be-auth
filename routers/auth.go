package routers

import (
	"gin-authflow/controllers"

	"github.com/gin-gonic/gin"
)

func authLogin(r *gin.RouterGroup) {
	r.POST("/login", controllers.AuthLogin)
	r.POST("/register", controllers.AuthLogin)
	r.POST("/forgot-password", controllers.ForgotPassword)
}

package routers

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	authLogin(r.Group("/auth"))
}

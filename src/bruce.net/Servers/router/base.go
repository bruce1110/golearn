package router

import (
	"github.com/gin-gonic/gin"
	"bruce.net/Servers/controllers"
)

func InitRoute(r *gin.Engine) {
	adminGroup := r.Group("admin")
	{
		adminGroup.GET("/test", controllers.Test)
	}
}

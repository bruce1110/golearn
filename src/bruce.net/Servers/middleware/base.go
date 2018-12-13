package middleware

import (
	"github.com/gin-gonic/gin"
)

func Base(c *gin.Context) {
	Auth(c)
}

package controllers

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	c.Abort()
	c.JSON(200, gin.H{
		"test": "ok",
	})
}
package middleware

import "github.com/gin-gonic/gin"

type params struct {
	Sign string `form:"sign"`
}

const key = "key123";

func Auth(c *gin.Context) {
	p := &params{}
	c.Bind(p)
	if (p.Sign != key) {
		c.Abort()
		c.JSON(401, gin.H{"status": "auth error"})
		return
	}
}

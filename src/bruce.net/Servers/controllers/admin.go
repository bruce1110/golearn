package controllers

import (
	"github.com/gin-gonic/gin"
	"bruce.net/Servers/model"
)

func Test(c *gin.Context) {
	p := &struct {
		Userid uint64 `form:"userid" json:"userid"`
	}{}
	if err := c.Bind(p); err != nil {
		c.Abort()
		c.JSON(401, gin.H{"err": "参数绑定失败"})
		return
	}
	if p.Userid == 0 {
		c.Abort()
		c.JSON(401, gin.H{"err": "请输入正确uid"})
		return
	}
	nickname, nick2, _ := model.GetOne(p.Userid)
	c.Abort()
	c.JSON(200, gin.H{
		"test":     "ok",
		"nickname": nickname,
		"nick2":    nick2,
	})
}

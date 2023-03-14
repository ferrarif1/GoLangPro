package service

import "github.com/gin-gonic/gin"
/*
处理业务逻辑
*/
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome",
	})
}

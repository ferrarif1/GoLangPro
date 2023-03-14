package service

import (
	"fmt"
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

/*
处理业务逻辑
*/
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	fmt.Println("GetUserList - userservice")
	fmt.Println(data)
	c.JSON(200, gin.H{
		"message": data,
	})
}

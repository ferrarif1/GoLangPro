package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

/*
处理url请求，从service返回给前端
*/

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	return r
}

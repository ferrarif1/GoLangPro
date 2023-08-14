package main

import (
	"net/http"
	"server/api/routers"
	"server/config"
	"server/core"
	"server/core/response"
	"server/middleware"
	"server/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// initRouter 初始化router
func initRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.Config.GinMode)
	router := gin.New()
	// 设置中间件
	router.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())
	// 特殊异常处理
	router.NoMethod(response.NoMethod)
	router.NoRoute(response.NoRoute)
	// 注册路由
	group := router.Group("/api")
	core.RegisterGroup(group, routers.WalletGroup, middleware.KeyAuth())
	core.RegisterGroup(group, routers.NftGroup, middleware.KeyAuth())
	return router
}

// initServer 初始化server
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		// 只接受本地请求，前半截可改成“127.0.0.1:”
		Addr:           ":" + strconv.Itoa(config.Config.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 初始化router
	router := initRouter()
	// 初始化验证器翻译
	util.ValidateUtil.Init()
	// 初始化server
	s := initServer(router)
	// 运行服务
	s.ListenAndServe().Error()
}

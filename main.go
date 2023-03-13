package main

import (
	"gin-mongodb-pack/bootstrap"
	"gin-mongodb-pack/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 载入配置
	bootstrap.InitializeConfig()
	if global.App.Config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode) // 在生产模式中，将 gin 设为 release 模式
	}

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log initial success.")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeMongoDB()
	global.App.Log.Info("mongoDB connect success.")
	defer func() { bootstrap.CloseMongoDB() }() // 程序结束时断开链接

	// 初始化 gin
	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// 同步启动期间产生的日志
	if err := global.App.Log.Sync(); err != nil {
		global.App.Log.Panic(err)
	}

	// 启动服务器，在8000端口
	_ = r.Run(":" + global.App.Config.Port)
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// 启动服务器，在8000端口
	_ = r.Run(":8000")

	// 打开浏览器访问 http://127.0.0.1:8000/ping
	// 理想结果：浏览器显示 'pong' 字符串
}

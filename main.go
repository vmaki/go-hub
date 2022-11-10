package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Hello World",
		})
	})

	// 处理 404 请求
	r.NoRoute(func(ctx *gin.Context) {
		accept := ctx.GetHeader("Accept")
		if strings.Contains(accept, "text/html") {
			ctx.String(http.StatusNotFound, "该页面不存在")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": "该接口不存在",
			})
		}
	})

	_ = r.Run(":7001")
}

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-hub/routes"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(r)

	// 注册 API 路由
	routes.RegisterAPIRouters(r)

	// 配置 404 路由
	setup404Handler(r)
}

func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(r *gin.Engine) {
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
}

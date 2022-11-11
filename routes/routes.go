package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRouters(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Hello World",
			})
		})

		v1.GET("/test/recovery", func(ctx *gin.Context) {
			panic("这是 panic 测试")

			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Hello Recovery",
			})
		})
	}
}

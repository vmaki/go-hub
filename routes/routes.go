package routes

import (
	"github.com/gin-gonic/gin"
	v12 "go-hub/app/http/controller/api/v1"
)

func RegisterAPIRouters(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		testGroup := v1.Group("/test")
		{
			api := v12.TestController{}
			testGroup.GET("/", api.Index)
			testGroup.GET("/recovery", api.Recovery)
			testGroup.GET("/db", api.Db)
			testGroup.GET("/redis", api.Redis)
			testGroup.POST("/vali", api.Vali)
		}
	}
}

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-hub/app/http/controller/api"
	"go-hub/app/service"
	"go-hub/pkg/redis"
	"go-hub/pkg/response"
	"time"
)

// TestController 测试接口专用控制器
type TestController struct {
	api.BaseAPIController
}

func (tc *TestController) Index(ctx *gin.Context) {
	response.Success(ctx, "Hello World!")
}

func (tc *TestController) Recovery(ctx *gin.Context) {
	panic("这是 panic 测试")

	// 正常情况是访问不到下面的代码
	response.Success(ctx)
}

func (tc TestController) Db(ctx *gin.Context) {
	userService := service.UserService{}
	res := userService.IsPhoneExist("15913395633")

	response.Success(ctx, "用户是否存在："+cast.ToString(res))
}

func (tc TestController) Redis(ctx *gin.Context) {
	res := redis.Client.Set("my-name", "VMaki", time.Duration(30)*time.Second)
	fmt.Println("是否写成功：", res)

	res2 := redis.Client.Get("my-name")
	response.Success(ctx, "缓存："+res2)
}

package v1

import (
	"github.com/gin-gonic/gin"
	"go-hub/app/http/controller/api"
	"go-hub/app/service/dto"
	"go-hub/pkg/jwt"
	"go-hub/pkg/request"
	"go-hub/pkg/response"
)

type AuthController struct {
	api.BaseAPIController
}

func (c AuthController) Login(ctx *gin.Context) {
	req := dto.LoginReq{}
	if ok := request.Validate(ctx, &req); !ok {
		return
	}

	if req.Phone != "15913395633" || req.Password != "123456" {
		response.Error(ctx, "账户或密码错误")
		return
	}

	token := jwt.NewJWT().CreateToken(1, "Maki")
	data := &dto.LoginResp{
		Token: token,
	}

	response.JSON(ctx, data)
}

func (c AuthController) RefreshToken(ctx *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(ctx)
	if err != nil {
		response.Error(ctx, err.Error())
	}

	data := &dto.RefreshTokenResp{
		Token: token,
	}
	response.JSON(ctx, data)
}

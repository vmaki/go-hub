package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"go-hub/pkg/response"
)

func Validate(ctx *gin.Context, req IRequest) bool {
	// 解析 JSON 请求
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.BadRequest(ctx)
		return false
	}

	// 表单验证
	errs := req.Generate(req)
	if errs != "" {
		response.BadRequest(ctx, errs)
		return false
	}

	return true
}

func GoValidate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	errs := govalidator.New(opts).ValidateStruct()
	if len(errs) > 0 {
		str := ""
		for _, v := range errs {
			str = v[0]
			break
		}

		return str
	}

	return ""
}

package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-hub/pkg/request"
)

type ValiReq struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func (s *ValiReq) Generate(data interface{}) string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return request.GoValidate(data, rules, messages)
}

type ValiResp struct {
	Name string
	Age  int
}

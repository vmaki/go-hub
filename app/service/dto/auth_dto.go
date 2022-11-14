package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-hub/pkg/request"
)

type LoginReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`
	Password string `json:"password,omitempty" valid:"password"`
}

func (s *LoginReq) Generate(data interface{}) string {
	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11"},
		"password": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	return request.GoValidate(data, rules, messages)
}

type LoginResp struct {
	Token string
}

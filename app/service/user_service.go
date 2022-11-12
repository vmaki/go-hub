package service

import (
	"go-hub/app/model"
	"go-hub/pkg/database"
)

type UserService struct {
	BaseService
}

func (s UserService) IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(model.UserModel{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

package controller

import (
	"permission-api/model"
	"permission-api/util"

	"go.mongodb.org/mongo-driver/bson"
)

type CreateUserOpts struct {
	UserId   string `json:"userId" binding:"required"`                   // 帳號
	Password string `json:"password" binding:"required"`                 // 密碼
	Name     string `json:"name" binding:"required"`                     // 姓名
	UserType string `json:"userType" binding:"required" example:"OTHER"` // 使用者類別 MANAGER=管理層, EMPLOYEE=員工, OTHER=其他, SYSTEM=系統
}

// 建立使用者
func CreateUser(opts CreateUserOpts, result *model.User) error {
	passwordSalt, err := util.GenerateHex(16)
	if err != nil {
		return err
	}
	passwordHash := util.HashPasswordWithSalt(opts.Password, passwordSalt)

	user := &model.User{
		UserId:       opts.UserId,
		PasswordSalt: passwordSalt,
		PasswordHash: passwordHash,
		Name:         opts.Name,
		UserType:     model.UserType(opts.UserType),
	}
	return model.Insert(model.UserCollName, user, result)
}

// 取得使用者
func GetUserByUserId(userId string, result *model.User) error {
	return model.Get(model.UserCollName, bson.D{{Key: "userId", Value: userId}}, &result)
}

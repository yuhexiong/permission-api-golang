package controller

import (
	"errors"
	"permission-api/controller/sessionController"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginOpts struct {
	UserId   string `json:"userId" binding:"required"`   // 帳號
	Password string `json:"password" binding:"required"` // 密碼
}

// 登入
func Login(opts LoginOpts) (*string, error) {
	// 取得使用者
	user := model.User{}
	if err := GetUser(opts.UserId, &user); err != nil {
		return nil, err
	}

	// 驗證密碼
	if !util.ValidatePassword(user.PasswordHash, user.PasswordSalt, opts.Password) {
		return nil, errors.New("")
	}

	var sessionsByUser []*model.Session
	if err := sessionController.FindSessionByUserId(user.ID, &sessionsByUser); err != nil {
		return nil, err
	}
	// 如果是已登入的系統使用者, 則不再登入
	if user.UserType == model.UserTypeSystem && len(sessionsByUser) > 0 {
		return nil, errors.New("system user already login")
	}
	// 移除目前有的登入憑證
	if len(sessionsByUser) > 0 {
		sessionController.DeleteSessionByUserId(user.ID)
	}

	// 產生 token
	var token *string
	if err := middleware.CreateToken(&user, token); err != nil {
		return nil, err
	}

	// 產生 session
	if err := sessionController.CreateSession(*token, &user); err != nil {
		return nil, err
	}

	return token, nil
}

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
func GetUser(userId string, result *model.User) error {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	return model.Get(model.UserCollName, bson.D{{Key: "_id", Value: objectId}}, &result)
}

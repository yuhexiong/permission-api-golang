package controller

import (
	"errors"
	"permission-api/controller/sessionController"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/util"
)

type LoginOpts struct {
	UserId   string `json:"userId" binding:"required"`   // 帳號
	Password string `json:"password" binding:"required"` // 密碼
}

// 登入
func Login(opts LoginOpts) (*string, error) {
	// 取得使用者
	user := model.User{}
	if err := GetUserByUserId(opts.UserId, &user); err != nil {
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
	var token string
	if err := middleware.CreateToken(&user, &token); err != nil {
		return nil, err
	}

	// 產生 session
	if err := sessionController.CreateSession(token, &user); err != nil {
		return nil, err
	}

	return &token, nil
}

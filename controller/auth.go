package controller

import (
	"errors"
	"os"
	"permission-api/controller/permissionController"
	"permission-api/controller/sessionController"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/util"
)

func InitAdminUser() {
	// 建立系統使用者
	adminUser := model.User{}
	GetUserByUserId("admin", &adminUser)

	// 如果沒有admin, 則新建一位
	if adminUser.ID == nil {
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		if adminPassword == "" {
			util.RedLog("should provide admin password")
		}

		createUserOpts := CreateUserOpts{
			UserId:   "admin",
			Password: adminPassword,
			Name:     "系統使用者",
			UserType: string(model.UserTypeSystem),
		}

		err := CreateUser(createUserOpts, &adminUser)
		if err != nil {
			panic(err)
		}
	}

	// 刪除 admin 的所有權限
	DeleteUserPermissionByUserOId(adminUser.ID)

	// 取得目前有的權限
	permissions := []*model.Permission{}
	if err := permissionController.FindPermission(permissionController.FindPermissionOpts{}, &permissions); err != nil {
		panic(err)
	}

	// 賦予系統使用者所有權限
	for _, permission := range permissions {
		createUserPermission := CreateUserPermissionOpts{
			UserOId:       adminUser.ID,
			PermissionOId: permission.ID,
			Operations:    []model.PermissionOp{model.READ, model.WRITE},
		}
		if err := CreateUserPermission(createUserPermission, nil); err != nil {
			panic(err)
		}
	}
}

type LoginOpts struct {
	UserId   string `json:"userId" binding:"required"`   // 帳號
	Password string `json:"password" binding:"required"` // 密碼
}

// 登入
func Login(opts LoginOpts) (*string, error) {
	// 取得使用者
	user := model.User{}
	if err := GetUserByUserId(opts.UserId, &user); err != nil {
		return nil, util.UserNotFoundError(err.Error())
	}

	// 驗證密碼
	if !util.ValidatePassword(user.PasswordHash, user.PasswordSalt, opts.Password) {
		return nil, util.WrongPasswordError("on login")
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
		sessionController.DeleteSessionByUserOId(user.ID)
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

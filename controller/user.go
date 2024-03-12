package controller

import (
	"permission-api/model"
	"permission-api/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordOpts struct {
	PasswordSalt string `bson:"_password_salt,omitempty" json:"-"`
	PasswordHash string `bson:"_password_hash,omitempty" json:"-"`
}

type ResetPasswordOpts struct {
	OldPassword string `json:"oldPassword" binding:"required"` // 舊密碼
	NewPassword string `json:"newPassword" binding:"required"` // 新密碼
}

// 修改自己密碼
func ResetPassword(user *model.User, opts ResetPasswordOpts) (bool, error) {
	// 驗證密碼
	if !util.ValidatePassword(user.PasswordHash, user.PasswordSalt, opts.OldPassword) {
		return false, util.WrongPasswordError("on reset password")
	}

	passwordSalt, err := util.GenerateHex(16)
	if err != nil {
		return false, err
	}

	passwordHash := util.HashPasswordWithSalt(opts.NewPassword, passwordSalt)

	if err = model.Update(
		model.UserCollName,
		user.ID,
		PasswordOpts{PasswordSalt: passwordSalt, PasswordHash: passwordHash}); err != nil {
		return false, err
	}

	return true, nil
}

// 修改他人密碼
func ChangePassword(user *model.User, password string) (bool, error) {
	passwordSalt, err := util.GenerateHex(16)
	if err != nil {
		return false, err
	}

	passwordHash := util.HashPasswordWithSalt(password, passwordSalt)

	if err = model.Update(
		model.UserCollName,
		user.ID,
		PasswordOpts{PasswordSalt: passwordSalt, PasswordHash: passwordHash}); err != nil {
		return false, err
	}

	return true, nil
}

type CreateUserOpts struct {
	UserId   string         `json:"userId" binding:"required" example:"admin"`         // 帳號
	Password string         `json:"password" binding:"required" example:"Abc12345678"` // 密碼
	Name     string         `json:"name" binding:"required" example:"系統使用者"`           // 姓名
	UserType model.UserType `json:"userType" binding:"required" example:"OTHER"`       // 使用者類別 MANAGER=管理層, EMPLOYEE=員工, OTHER=其他, SYSTEM=系統
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

type FindUserOpts struct {
	UserId   *string         `json:"userId" example:"admin"`    // 帳號(支援正規表達式)
	Name     *string         `json:"name" example:"系統使用者"`      // 姓名(支援正規表達式)
	UserType *model.UserType `json:"userType"  example:"OTHER"` // 使用者類別 MANAGER=管理層, EMPLOYEE=員工, OTHER=其他, SYSTEM=系統
}

// 搜尋使用者
func FindUser(opts FindUserOpts, result *[]*model.User) error {
	filter := bson.D{}

	if opts.UserId != nil {
		filter = append(filter, bson.E{Key: "userId", Value: primitive.Regex{Pattern: *opts.UserId, Options: ""}})
	}

	if opts.Name != nil {
		filter = append(filter, bson.E{Key: "name", Value: primitive.Regex{Pattern: *opts.Name, Options: ""}})
	}

	if opts.UserType != nil {
		filter = append(filter, bson.E{Key: "userType", Value: opts.UserType})
	}

	return model.Find(model.UserCollName, filter, result)
}

// 依userOId取得使用者
func GetUserByUserOId(objectId *primitive.ObjectID, result *model.User) error {
	return model.Get(model.UserCollName, objectId, &result)
}

// 依userId取得使用者
func GetUserByUserId(userId string, result *model.User) error {
	return model.GetByFilter(model.UserCollName, bson.D{{Key: "userId", Value: userId}}, &result)
}

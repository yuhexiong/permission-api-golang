package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserCollName = "user"

type UserType string

const (
	UserTypeManager  UserType = "MANAGER"
	UserTypeEmployee UserType = "EMPLOYEE"
	UserTypeOther    UserType = "OTHER"
	UserTypeSystem   UserType = "SYSTEM"
)

type User struct {
	ID           *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"`
	BaseData     `bson:"inline"`
	UserId       string   `bson:"userId" json:"userId" example:"abd1234"` // 帳號
	PasswordSalt string   `bson:"_password_salt,omitempty" json:"-"`
	PasswordHash string   `bson:"_password_hash,omitempty" json:"-"`
	Name         string   `bson:"name" json:"name" example:"欸逼西滴"`          // 姓名
	UserType     UserType `bson:"userType" json:"userType" example:"OTHER"` // 使用者類別 MANAGER=管理層, EMPLOYEE=員工, OTHER=其他, SYSTEM=系統
}

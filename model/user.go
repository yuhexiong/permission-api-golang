package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserCollName = "user"

type UserTypeOpt string

const (
	UserTypeManager  UserTypeOpt = "MANAGER"
	UserTypeEmployee UserTypeOpt = "EMPLOYEE"
	UserTypeOther    UserTypeOpt = "OTHER"
	UserTypeSystem   UserTypeOpt = "SYSTEM"
)

type User struct {
	ID           *primitive.ObjectID `bson:"id,omitempty" json:"id" example:"623853b9503ce2ecdd221c94"`
	Status       *uint8              `bson:"status,omitempty" json:"status" example:"0"` // 0: 正常, 9: 刪除
	CreatedAt    *time.Time          `bson:"createdAt,omitempty" json:"createdAt" example:"2022-03-21T10:30:17.711Z"`
	UpdatedAt    *time.Time          `bson:"updatedAt,omitempty" json:"updatedAt" example:"2022-03-21T10:30:17.711Z"`
	PasswordSalt string              `bson:"_password_salt,omitempty" json:"-"`
	PasswordHash string              `bson:"_password_hash,omitempty" json:"-"`
	Name         string              `bson:"name,omitempty" json:"name" example:"艾瑞克"`           // 姓名
	UserType     string              `bson:"userType,omitempty" json:"userType" example:"OTHER"` // 使用者類別 MANAGER=管理層, EMPLOYEE=員工, OTHER=其他, SYSTEM=系統
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PermissionCollName = "permission"

type Permission struct {
	ID        *primitive.ObjectID `bson:"id,omitempty" json:"id" example:"623853b9503ce2ecdd221c94"`
	Status    *uint8              `bson:"status,omitempty" json:"status" example:"0"` // 0: 正常, 9: 刪除
	CreatedAt *time.Time          `bson:"createdAt,omitempty" json:"createdAt" example:"2022-03-21T10:30:17.711Z"`
	UpdatedAt *time.Time          `bson:"updatedAt,omitempty" json:"updatedAt" example:"2022-03-21T10:30:17.711Z"`
	Category  string              `bson:"category,omitempty" json:"category" example:"USER"` // 權限種類
	Code      string              `bson:"code,omitempty" json:"code" example:"createUser"`   // 權限代號
}

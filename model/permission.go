package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PermissionCollName = "permission"

type Permission struct {
	ID       *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"`
	BaseData `bson:"inline"`
	Category string `bson:"category" json:"category" example:"USER"` // 權限種類
	Code     string `bson:"code" json:"code" example:"createUser"`   // 權限代號
}

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MapUserPermissionCollName = "mapUserPermission"

// 讀、寫兩種執行權限
type PermissionOp string

const (
	READ  PermissionOp = "R"
	WRITE PermissionOp = "W"
)

type MapUserPermission struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"`
	UserOId       *primitive.ObjectID `bson:"userOId" json:"userOId"`                               // 使用者 objectId
	PermissionOId *primitive.ObjectID `bson:"permissionOId" json:"permissionOId"`                   // 權限 objectId
	Operations    []PermissionOp      `bson:"operations,omitempty" json:"operations" example:"W,R"` // 讀或寫
}

package controller

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserPermissionOpts struct {
	UserOId       *primitive.ObjectID  `bson:"userOId" json:"userOId"`                     // 使用者 objectId
	PermissionOId *primitive.ObjectID  `bson:"permissionOId" json:"permissionOId"`         // 權限 objectId
	Operations    []model.PermissionOp `bson:"operations" json:"operations" example:"W,R"` // 讀或寫
}

// 建立使用者與權限對應關係
func CreateUserPermission(opts CreateUserPermissionOpts, result *model.MapUserPermission) error {
	userPermission := &model.MapUserPermission{
		UserOId:       opts.UserOId,
		PermissionOId: opts.PermissionOId,
		Operations:    opts.Operations,
	}
	return model.Insert(model.MapUserPermissionCollName, userPermission, result)
}

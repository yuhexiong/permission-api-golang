package permissionController

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindPermissionOpts struct {
	Category *string `json:"category" bson:"category,omitempty" example:"USER"` // 權限種類
	Code     *string `json:"code" bson:"code,omitempty" example:"createUser"`   // 權限代號
}

// 取得權限
func FindPermission(opts FindPermissionOpts, result *[]*model.Permission) error {
	filter := bson.D{}

	if opts.Category != nil {
		filter = append(filter, bson.E{Key: "category", Value: opts.Category})
	}

	if opts.Code != nil {
		filter = append(filter, bson.E{Key: "code", Value: opts.Code})
	}

	return model.Find(model.PermissionCollName, filter, &result)
}

type CreatePermissionOpts struct {
	Category string `json:"category,omitempty" example:"USER"`                      // 權限種類
	Code     string `json:"code,omitempty" binding:"required" example:"createUser"` // 權限代號
}

// 新增權限
func CreatePermission(opts CreatePermissionOpts, result *model.Permission) error {
	permission := model.Permission{
		Category: opts.Category,
		Code:     opts.Code,
	}

	return model.Insert(model.PermissionCollName, permission, result)
}

// 啟用權限
func EnablePermission(objectId *primitive.ObjectID) error {
	return model.Enable(model.PermissionCollName, objectId)
}

// 刪除權限
func DeletePermission(objectId *primitive.ObjectID) error {
	return model.Delete(model.PermissionCollName, objectId, false)
}

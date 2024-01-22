package permissionController

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindPermissionOptions struct {
	Category *string `json:"category" bson:"category,omitempty" example:"USER"` // 權限種類
	Code     *string `json:"code" bson:"code,omitempty" example:"createUser"`   // 權限代號
}

// 取得所有權限
func FindPermission(opts FindPermissionOptions, result *[]*model.Permission) error {
	filter := bson.D{}

	if opts.Category != nil {
		filter = append(filter, bson.E{Key: "category", Value: opts.Category})
	}

	if opts.Code != nil {
		filter = append(filter, bson.E{Key: "code", Value: opts.Code})
	}

	return model.Find(model.PermissionCollName, filter, &result)
}

type CreatePermissionOptions struct {
	Category string `json:"category,omitempty" example:"USER"`                      // 權限種類
	Code     string `json:"code,omitempty" binding:"required" example:"createUser"` // 權限代號
}

func CreatePermission(opts CreatePermissionOptions, result *model.Permission) error {
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

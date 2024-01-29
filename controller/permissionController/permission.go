package permissionController

import (
	"fmt"
	"permission-api/model"
	"permission-api/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	return model.Find(model.PermissionCollName, filter, result)
}

type CreatePermissionOpts struct {
	Category string `json:"category,omitempty" example:"USER"`                      // 權限種類
	Code     string `json:"code,omitempty" binding:"required" example:"createUser"` // 權限代號
}

// 新增權限
func CreatePermission(opts CreatePermissionOpts, result *model.Permission) error {
	permission := model.Permission{
		Status:   util.GetPointer(model.NormalStatus),
		Category: opts.Category,
		Code:     opts.Code,
	}

	return model.Insert(model.PermissionCollName, &permission, result)
}

// 啟用權限
func EnablePermission(objectId *primitive.ObjectID) error {
	return model.Enable(model.PermissionCollName, objectId)
}

// 刪除權限
func DeletePermission(objectId *primitive.ObjectID) error {
	return model.Delete(model.PermissionCollName, objectId, false)
}

type PermissionDetails struct {
	Category string `bson:"category,omitempty" json:"category" example:"USER"`
	Code     string `bson:"code,omitempty" json:"code" example:"createUser"`
}

type MapUserPermissionWithDetails struct {
	model.MapUserPermission `bson:",inline"`
	PermissionDetails       `bson:",inline"`
}

type PermissionInfo struct {
	UserOId       string                           `json:"userOId" example:"623853b9503ce2ecdd221c94"` // 始俑者 ObjectId
	PermissionMap *map[string][]model.PermissionOp `json:"-"`                                          // 權限key: [category]-[code] value: "R,W"
}

// 取得該使用者的權限
func GetPermissionInfoByUser(userOId *primitive.ObjectID) (*map[string][]model.PermissionOp, error) {
	var userPermissions []*MapUserPermissionWithDetails

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "userOId", Value: userOId}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: model.PermissionCollName},
			{Key: "localField", Value: "permissionOId"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "permissionDetails"},
		}}},
		{{Key: "$unwind", Value: "$permissionDetails"}},
	}

	err := model.FindByPipeline(model.MapUserPermissionCollName, pipeline, &userPermissions)
	if err != nil {
		return nil, err
	}

	permissionOp := make(map[string][]model.PermissionOp)
	for _, userPermission := range userPermissions {
		permissionKey := fmt.Sprintf("%s-%s", userPermission.PermissionDetails.Category, userPermission.PermissionDetails.Code)
		permissionOp[permissionKey] = userPermission.Operations
	}

	return &permissionOp, nil
}

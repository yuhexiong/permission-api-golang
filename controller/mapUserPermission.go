package controller

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateUserPermissionOpts struct {
	UserOId       *primitive.ObjectID  `bson:"userOId" json:"userOId" binding:"required"`                                  // 使用者 objectId
	PermissionOId *primitive.ObjectID  `bson:"permissionOId" json:"permissionOId" binding:"required"`                      // 權限 objectId
	Operations    []model.PermissionOp `bson:"operations" json:"operations" binding:"required,permissionOp" example:"W,R"` // 讀或寫
}

// 建立使用者與權限對應關係
func CreateUserPermission(opts CreateUserPermissionOpts, result *model.MapUserPermission) error {
	userPermission := model.MapUserPermission{}
	if err := model.GetByFilter(model.MapUserPermissionCollName,
		bson.D{{Key: "userOId", Value: opts.UserOId}, {Key: "permissionOId", Value: opts.PermissionOId}},
		&userPermission); err != nil {
		// 沒有權限就新增
		if err == mongo.ErrNoDocuments {
			createdUserPermission := model.MapUserPermission{
				UserOId:       opts.UserOId,
				PermissionOId: opts.PermissionOId,
				Operations:    opts.Operations,
			}
			return model.Insert(model.MapUserPermissionCollName, &createdUserPermission, result)
		}

		return err
	} else {
		// 有權限就就更新並啟用
		return model.Update(model.MapUserPermissionCollName,
			userPermission.ID,
			bson.D{{Key: "operations", Value: opts.Operations}, {Key: "status", Value: model.NormalStatus}})
	}
}

type FindUserPermissionOpts struct {
	UserOId       *primitive.ObjectID `bson:"userOId" json:"userOId"`             // 使用者 objectId
	PermissionOId *primitive.ObjectID `bson:"permissionOId" json:"permissionOId"` // 權限 objectId
}

// 取得權限
func FindUserPermission(opts FindUserPermissionOpts, result *[]*model.MapUserPermission) error {
	filter := bson.D{}

	if opts.UserOId != nil {
		filter = append(filter, bson.E{Key: "userOId", Value: opts.UserOId})
	}

	if opts.PermissionOId != nil {
		filter = append(filter, bson.E{Key: "permissionOId", Value: opts.PermissionOId})
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: model.PermissionCollName},
			{Key: "localField", Value: "permissionOId"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: model.PermissionCollName},
		}}},
		{{Key: "$unwind", Value: "$permission"}},
	}

	return model.FindByPipeline(model.MapUserPermissionCollName, pipeline, result)
}

// 刪除使用者與權限對應關係
func DeleteUserPermission(objectId *primitive.ObjectID) error {
	return model.Delete(model.MapUserPermissionCollName, objectId, false)
}

// 刪除特定使用者的使用者與權限對應關係
func DeleteUserPermissionByUserOId(userOId *primitive.ObjectID) error {
	return model.DeleteByFilter(model.MapUserPermissionCollName, bson.D{{Key: "userOId", Value: userOId}}, true)
}

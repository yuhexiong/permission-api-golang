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
	userPermission := model.MapUserPermission{
		UserOId:       opts.UserOId,
		PermissionOId: opts.PermissionOId,
		Operations:    opts.Operations,
	}
	return model.Insert(model.MapUserPermissionCollName, &userPermission, result)
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
func DeleteUserPermission(userOId *primitive.ObjectID) error {
	return model.DeleteByFilter(model.MapUserPermissionCollName, bson.D{{Key: "userOId", Value: userOId}}, true)
}

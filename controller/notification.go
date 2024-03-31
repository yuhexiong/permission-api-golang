package controller

import (
	"errors"
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type createNotificationOpts struct {
	ToUserOId *primitive.ObjectID `bson:"toUserOId" json:"toUserOId" binding:"required" example:"abd1234"` // 接收者帳號id
	Content   string              `bson:"content" json:"content" binding:"required" example:"您有一個新的任務"`    // 通知內容
}

// 建立通知
func createNotification(opts createNotificationOpts, result *model.Notification) error {
	createdNotification := &model.Notification{
		ToUserOId: opts.ToUserOId,
		Content:   opts.Content,
	}
	return model.Insert(model.NotificationCollName, createdNotification, result)
}

type FindNotificationOpts struct {
	ToUserOId *primitive.ObjectID `json:"toUserOId" bson:"toUserOId" binding:"required" example:"623850247ea4cca15cd55303"` // 被發送帳號id
}

// 取得通知
func FindNotification(opts FindNotificationOpts, result *[]*model.Notification) error {
	filter := bson.D{{Key: "toUserOId", Value: opts.ToUserOId}}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sort", Value: bson.D{{Key: "createdAt", Value: -1}}}},
	}

	return model.FindByPipeline(model.NotificationCollName, pipeline, result)
}

// 已讀通知
func ReadNotification(objectId *primitive.ObjectID, userOId *primitive.ObjectID) error {
	notification := model.Notification{}
	if err := model.Get(model.NotificationCollName, objectId, &notification); err != nil {
		return err
	}

	if *notification.ToUserOId != *userOId {
		return errors.New("notification not sent to this user")
	}

	return model.Update(model.NotificationCollName, objectId, bson.D{{Key: "read", Value: true}})
}

// 已讀所有通知
func ReadAllNotification(userOId *primitive.ObjectID) error {
	return model.UpdateByFilter(model.NotificationCollName, bson.D{{Key: "toUserOId", Value: userOId}}, bson.D{{Key: "read", Value: true}})
}

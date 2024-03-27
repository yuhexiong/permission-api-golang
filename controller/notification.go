package controller

import (
	"errors"
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createNotificationOpts struct {
	ToUserOId *primitive.ObjectID `bson:"ToUserOId" json:"ToUserOId" binding:"required" example:"abd1234"` // 接收者帳號id
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

package controller

import (
	"permission-api/model"

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

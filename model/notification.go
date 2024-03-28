package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const NotificationCollName = "notification"

type Notification struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c95"`
	BaseData  `bson:"inline"`
	ToUserOId *primitive.ObjectID `bson:"toUserOId" json:"toUserOId" example:"abd1234"` // 接收者帳號id
	ToUser    *User               `bson:"toUser,omitempty" json:"toUser"`               // 接收者帳號
	Content   string              `bson:"content" json:"content" example:"您有一個新的任務"`    // 通知內容
	Read      *bool               `bson:"read,omitempty" json:"read" example:"true"`    // 已讀與否
}

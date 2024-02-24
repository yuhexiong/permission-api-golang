package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TaskCollName = "task"

type PriorityType string

const (
	PriorityTypeHigh   PriorityType = "HIGH"
	PriorityTypeMedium PriorityType = "MEDIUM"
	PriorityTypeLow    PriorityType = "LOW"
)

type Task struct {
	ID           *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"`
	BaseData     `bson:"inline"`
	FromUserOId  *primitive.ObjectID `bson:"fromUserOId" json:"fromUserOId" example:"abd1234"`       // 指派帳號id
	FromUser     *User               `bson:"fromUser,omitempty" json:"fromUser"`                     // 指派帳號
	ToUserOId    *primitive.ObjectID `bson:"ToUserOId" json:"ToUserOId" example:"abd1234"`           // 被指派帳號id
	ToUser       *User               `bson:"toUser,omitempty" json:"toUser"`                         // 指派帳號
	Deadline     string              `bson:"deadline" json:"deadline" example:"2024-01-03"`          // 結束日期
	PriorityType PriorityType        `bson:"priorityType" json:"priorityType" example:"HIGH"`        // 任務優先度 HIGH=高, MEDIUM=中, LOW=低
	Title        string              `bson:"title" json:"title" example:"完成表格"`                      // 標題
	Content      string              `bson:"content,omitempty" json:"content" example:"將1到10欄的內容補齊"` // 內容
	Checked      *bool               `bson:"checked,omitempty" json:"checked" example:"true"`        // 驗收
}

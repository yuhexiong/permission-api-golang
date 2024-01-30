package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SessionCollName = "session"

type SessionType string

// 一般使用者 與 系統使用者
const SessionTypeNormal SessionType = "normal"
const SessionTypeSystem SessionType = "system"

type Session struct {
	ID           *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"`
	BaseData     `bson:"inline"`
	UserOId      *primitive.ObjectID `bson:"userOId" json:"userOId"` // 使用者 objectId
	SessionToken string              `bson:"sessionToken"`
	Type         SessionType         `bson:"type"`
	ExpiresAt    *time.Time          `bson:"expiresAt,omitempty"`
}

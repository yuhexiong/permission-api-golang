package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SettingCollName = "setting"

type Setting struct {
	ID       *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"`
	BaseData `bson:"inline"`
	Code     string `bson:"code" json:"code" example:"timeFormat"`
	Value    string `bson:"value" json:"value" example:"YYYY-MM-DD HH:mm:ss"`
}

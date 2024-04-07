package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const LogCollName = "log"

type Log struct {
	Method       string              `json:"method" bson:"method"`
	URL          string              `json:"url" bson:"url"`
	Headers      interface{}         `json:"headers" bson:"headers"`
	UserOID      *primitive.ObjectID `json:"userOId" bson:"userOId"`
	StatusCode   int                 `json:"statusCode" bson:"statusCode"`
	RequestBody  interface{}         `json:"requestBody" bson:"requestBody"`
	ErrorCode    *string             `json:"errorCode" bson:"errorCode"`
	ErrorMessage *string             `json:"errorMessage" bson:"errorMessage"`
	CreatedAt    time.Time           `json:"createdAt" bson:"createdAt"`
}

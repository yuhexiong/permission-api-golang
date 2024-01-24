package controller

import (
	"permission-api/model"
	"permission-api/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 取得該使用者的的登入憑證
func FindSessionByUserId(objectId *primitive.ObjectID, result *[]*model.Session) error {
	return model.Find(model.SessionCollName, bson.D{{Key: "_id", Value: objectId}}, &result)
}

// 新增登入憑證
func CreateSession(token string, user *model.User) error {
	var expiresAt time.Time
	sessionType := model.SessionTypeNormal
	if model.UserType(user.UserType) == model.UserTypeSystem {
		expiresAt = time.Now().Add(util.SystemTokenLifeTime)
		sessionType = model.SessionTypeSystem
	} else {
		expiresAt = time.Now().Add(util.NormalTokenLifeTime)
	}

	session := model.Session{
		UserOId:      user.ID,
		SessionToken: token,
		Type:         sessionType,
		ExpiresAt:    &expiresAt,
	}

	return model.Insert(model.SessionCollName, session, nil)
}

// 取得該使用者的的登入憑證
func DeleteSessionByUserId(objectId *primitive.ObjectID, result *[]*model.Session) error {
	return model.Find(model.SessionCollName, bson.D{{Key: "_id", Value: objectId}}, &result)
}

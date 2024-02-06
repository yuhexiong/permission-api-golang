package sessionController

import (
	"permission-api/model"
	"permission-api/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 取得該使用者的的登入憑證
func FindSessionByUserId(objectId *primitive.ObjectID, result *[]*model.Session) error {
	return model.Find(model.SessionCollName, bson.D{{Key: "userOId", Value: objectId}}, result)
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

	return model.Insert(model.SessionCollName, &session, nil)
}

// 刪除該使用者的的登入憑證
func DeleteSessionByUserOId(userOId *primitive.ObjectID) error {
	return model.DeleteByFilter(model.SessionCollName, bson.D{{Key: "userOId", Value: userOId}}, true)
}

// 由token取得登入憑證
func GetSessionByToken(token string, result *model.Session) error {
	return model.GetByFilter(model.SessionCollName, bson.D{{Key: "sessionToken", Value: token}}, &result)
}

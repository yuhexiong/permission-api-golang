package controller

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
)

// 依code取得設定
func GetSettingByCode(code *string, result *model.User) error {
	return model.GetByFilter(model.SettingCollName, bson.D{{Key: "code", Value: *code}}, &result)
}

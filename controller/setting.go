package controller

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 依code取得設定
func GetSettingByCode(code *string, result *model.Setting) error {
	return model.GetByFilter(model.SettingCollName, bson.D{{Key: "code", Value: *code}}, &result)
}

// 更新設定值
func UpdateSettingValue(objectId *primitive.ObjectID, value *string) error {
	return model.Update(model.SettingCollName, objectId, bson.D{{Key: "value", Value: *value}})
}

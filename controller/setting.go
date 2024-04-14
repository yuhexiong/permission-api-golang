package controller

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
)

// 依code取得設定
func GetSettingByCode(code *string, result *model.Setting) error {
	return model.GetByFilter(model.SettingCollName, bson.D{{Key: "code", Value: *code}}, &result)
}

// 更新設定值
func UpdateSettingValue(code *string, value *string) error {
	setting := model.Setting{}
	err := GetSettingByCode(code, &setting)
	if err != nil {
		return err
	}

	return model.Update(model.SettingCollName, setting.ID, bson.D{{Key: "value", Value: *value}})
}

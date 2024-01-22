package permissionController

import (
	"permission-api/model"

	"go.mongodb.org/mongo-driver/bson"
)

type FindPermissionOptions struct {
	Category *string `json:"category" bson:"category,omitempty" example:"USER"` // 權限種類
	Code     *string `json:"code" bson:"code,omitempty" example:"createUser"`   // 權限代號
}

// 取得所有權限
func FindPermission(opts FindPermissionOptions, result *[]*model.Permission) error {
	filter := bson.D{}

	if opts.Category != nil {
		filter = append(filter, bson.E{Key: "category", Value: opts.Category})
	}

	if opts.Code != nil {
		filter = append(filter, bson.E{Key: "code", Value: opts.Code})
	}

	if err := model.Find(model.PermissionCollName, filter, &result); err != nil {
		return err
	}

	return nil
}

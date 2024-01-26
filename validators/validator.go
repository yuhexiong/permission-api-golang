package validators

import (
	"permission-api/model"
	"permission-api/util"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("permissionOp", OperationsValidation)
	} else {
		util.RedLog("init validator error")
	}
}

func OperationsValidation(fl validator.FieldLevel) bool {
	operations := fl.Field().Interface().([]model.PermissionOp)

	if len(operations) == 0 {
		return true
	}

	for _, operation := range operations {
		switch operation {
		case
			model.READ,
			model.WRITE:
			continue
		default:
			return false
		}
	}

	return true
}

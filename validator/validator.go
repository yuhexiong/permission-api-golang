package validator

import (
	"permission-api/model"
	"permission-api/util"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("userType", UserTypeValidation); err != nil {
			util.RedLog("init userType validator error")
		}
		if err := v.RegisterValidation("permissionOp", OperationsValidation); err != nil {
			util.RedLog("init permissionOp validator error")
		}

	} else {
		util.RedLog("init validator error")
	}
}

func UserTypeValidation(fl validator.FieldLevel) bool {
	userType := fl.Field().Interface().(model.UserType)

	switch userType {
	case
		model.UserTypeManager,
		model.UserTypeEmployee,
		model.UserTypeOther,
		model.UserTypeSystem:
		break
	default:
		return false
	}

	return true
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

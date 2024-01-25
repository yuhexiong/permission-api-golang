package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"permission-api/controller/permissionController"
	"permission-api/model"

	"github.com/gin-gonic/gin"
)

// 驗證權限 middleware
func PermissionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		PermsDef, PermsOp := permissionController.GetApiPermission(c.FullPath(), c.Request.Method)

		// 無需驗證權限
		if PermsDef == nil {
			return
		}

		// 檢查使用者是否擁有指定的權限
		permissionMap := GetPermissionMap(c)
		hasPermission := CheckPermission(permissionMap, *PermsDef, *PermsOp)
		if !hasPermission {
			c.AbortWithError(http.StatusBadRequest, errors.New(""))
		}
		c.Next()
	}
}

// 驗證是否有此權限
func CheckPermission(permissionMap *map[string][]model.PermissionOp, pDef permissionController.PermissionDef, ops model.PermissionOp) bool {
	permissionKey := fmt.Sprintf("%s-%s", pDef.Category, pDef.Code)
	permissionOps := (*permissionMap)[permissionKey]

	for _, op := range permissionOps {
		if op == ops {
			return true
		}
	}

	return false
}

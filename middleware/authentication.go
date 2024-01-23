package middleware

import (
	"permission-api/controller/permissionController"

	"github.com/gin-gonic/gin"
)

type AccountInfo struct {
	UserOId       string                                          `json:"userOId" example:"623853b9503ce2ecdd221c94"` // 始俑者 ObjectId
	PermissionMap *map[string][]permissionController.PermissionOp `json:"-"`                                          // 權限key: [category]-[code] value: "R","W"
}

func SetAccountInfo(c *gin.Context, accountInfo *AccountInfo) {
	if accountInfo != nil {
		c.Set("accountInfo", accountInfo)
	}
}

func GetAccountInfo(c *gin.Context) *AccountInfo {
	accountInfo, exists := c.Get("accountInfo")

	if !exists {
		return nil
	}

	return accountInfo.(*AccountInfo)
}

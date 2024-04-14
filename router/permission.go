package router

import (
	"net/http"
	"permission-api/controller/permissionController"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
)

func InitPermissionRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/find", findPermission)
}

func findPermission(c *gin.Context) {
	var opts permissionController.FindPermissionOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	permissions := []*model.Permission{}
	if err := permissionController.FindPermission(opts, &permissions); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, permissions)
	c.Next()
}

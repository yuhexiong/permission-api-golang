package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
)

func InitUserPermissionRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/", createUserPermission)
	RouterPerms(routerGroup, http.MethodPost, "/find", findUserPermission)
}

func createUserPermission(c *gin.Context) {
	var createUserPermissionOpts controller.CreateUserPermissionOpts

	if err := c.ShouldBindJSON(&createUserPermissionOpts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	userPermission := model.MapUserPermission{}
	err := controller.CreateUserPermission(createUserPermissionOpts, &userPermission)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, userPermission)
}

func findUserPermission(c *gin.Context) {
	var opts controller.FindUserPermissionOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	userPermissions := []*model.MapUserPermission{}
	if err := controller.FindUserPermission(opts, &userPermissions); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, userPermissions)
}

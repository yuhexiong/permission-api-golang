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

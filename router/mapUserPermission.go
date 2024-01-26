package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/model"
	"permission-api/response"

	"github.com/gin-gonic/gin"
)

func InitUserPermissionRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/", createUserPermission)
}

func createUserPermission(c *gin.Context) {
	var createUserPermissionOpts controller.CreateUserPermissionOpts

	if err := c.ShouldBindJSON(&createUserPermissionOpts); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	userPermission := model.MapUserPermission{}
	err := controller.CreateUserPermission(createUserPermissionOpts, &userPermission)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	response.ResFormat(c, http.StatusOK, 0, userPermission)
}

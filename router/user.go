package router

import (
	"errors"
	"net/http"
	"permission-api/controller"
	"permission-api/controller/permissionController"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "", createUser)
}

func createUser(c *gin.Context) {
	var createUserOpts controller.CreateUserOpts

	if err := c.ShouldBindJSON(&createUserOpts); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 如果希望建立系統使用者, 則要有建立系統使用者的權限
	if model.UserType(createUserOpts.UserType) == model.UserTypeSystem {
		permissionMap := middleware.GetPermissionMap(c)
		permissionDef := permissionController.PermissionsMap[strings.ToLower("CreateSystemAccount")]
		if !middleware.CheckPermission(permissionMap, permissionDef, model.WRITE) {
			c.AbortWithError(http.StatusBadRequest, errors.New("not have permission to create system account"))
			return
		}
	}

	createdUser := model.User{}
	err := controller.CreateUser(createUserOpts, &createdUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response.ResFormat(c, http.StatusOK, 0, createdUser)
}

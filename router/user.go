package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/controller/permissionController"
	"permission-api/controller/sessionController"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/logout", logout)
	RouterPerms(routerGroup, http.MethodPost, "", createUser)
	RouterPerms(routerGroup, http.MethodPost, "/find", findUser)
}

func logout(c *gin.Context) {
	userOId := middleware.GetUserOId(c)
	// 取得使用者
	user := model.User{}
	if err := controller.GetUserByUserOId(userOId, &user); err != nil {
		response.AbortError(c, util.UserNotFoundError(err.Error()))
	}

	// 系統使用者永遠不登出
	if user.UserType != model.UserTypeSystem {
		sessionController.DeleteSessionByUserOId(userOId)
	}

	response.SuccessFormat(c, gin.H{})
}

func createUser(c *gin.Context) {
	var createUserOpts controller.CreateUserOpts

	if err := c.ShouldBindJSON(&createUserOpts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	// 如果希望建立系統使用者, 則要有建立系統使用者的權限
	if model.UserType(createUserOpts.UserType) == model.UserTypeSystem {
		permissionMap := middleware.GetPermissionMap(c)
		permissionDef := permissionController.PermissionsMap[strings.ToLower("CreateSystemAccount")]
		if !middleware.CheckPermission(permissionMap, permissionDef, model.WRITE) {
			response.AbortError(c, util.PermissionDeniedError("on create system user"))
			return
		}
	}

	createdUser := model.User{}
	err := controller.CreateUser(createUserOpts, &createdUser)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, createdUser)
}

func findUser(c *gin.Context) {
	var opts controller.FindUserOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	users := []*model.User{}
	if err := controller.FindUser(opts, &users); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, users)
}

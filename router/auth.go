package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/controller/sessionController"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
)

func InitUnAuthRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/login", login)
}

func InitAuthRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/logout", logout)
}

func login(c *gin.Context) {
	var loginOpt controller.LoginOpts

	if err := c.ShouldBindJSON(&loginOpt); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	token, err := controller.Login(loginOpt)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, token)
}

func logout(c *gin.Context) {
	userOId := middleware.GetUserOId(c)
	// 取得使用者
	user := model.User{}
	if err := controller.GetUserByUserOId(userOId, &user); err != nil {
		response.AbortError(c, util.UserNotFoundError(err.Error()))
		return
	}

	// 系統使用者永遠不登出
	if user.UserType != model.UserTypeSystem {
		sessionController.DeleteSessionByUserOId(userOId)
	}

	response.SuccessFormat(c, gin.H{})
}

package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/login", login)
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

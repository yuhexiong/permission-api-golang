package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/response"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/login", login)
}

func login(c *gin.Context) {
	var loginOpt controller.LoginOpts

	if err := c.ShouldBindJSON(&loginOpt); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	token, err := controller.Login(loginOpt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	response.ResFormat(c, http.StatusOK, 0, token)
}

package router

import (
	"errors"
	"net/http"
	"permission-api/controller"
	"permission-api/response"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/login", login)
}

func login(c *gin.Context) {
	var loginOpt controller.LoginOpts

	if err := c.ShouldBindJSON(&loginOpt); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
	}

	token, err := controller.Login(loginOpt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid token"))
	}

	response.ResFormat(c, http.StatusOK, 0, token)
}

package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/model"
	"permission-api/response"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/", createUser)
}

func createUser(c *gin.Context) {
	var createUserOpts controller.CreateUserOpts

	if err := c.ShouldBindJSON(&createUserOpts); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	user := model.User{}
	err := controller.CreateUser(createUserOpts, &user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	response.ResFormat(c, http.StatusOK, 0, user)
}

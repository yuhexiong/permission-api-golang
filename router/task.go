package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
)

func InitTaskRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "", createTask)
}

func createTask(c *gin.Context) {
	var opts controller.CreateTaskOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	userOId := middleware.GetUserOId(c)

	task := model.Task{}
	if err := controller.CreateTask(opts, userOId, &task); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, task)
}

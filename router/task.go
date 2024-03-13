package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitTaskRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "", createTask)
	RouterPerms(routerGroup, http.MethodPost, "/find", findTask)
	RouterPerms(routerGroup, http.MethodPatch, "/:id/checked/:checked", checkTask)
	RouterPerms(routerGroup, http.MethodDelete, "/:id", deleteTask)
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

func findTask(c *gin.Context) {
	var opts controller.FindTaskOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	tasks := []*model.Task{}
	if err := controller.FindTask(opts, &tasks); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, tasks)
}

type checkedTaskReqParm struct {
	ID      string `uri:"id" binding:"required"`
	Checked bool   `uri:"checked" binding:"required"`
}

func checkTask(c *gin.Context) {
	var params checkedTaskReqParm
	if err := c.ShouldBindUri(&params); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	objectId, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	userOId := middleware.GetUserOId(c)

	if err := controller.CheckTask(&objectId, userOId, &params.Checked); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, gin.H{})
}

type deleteTaskReqParm struct {
	ID string `uri:"id" binding:"required"`
}

func deleteTask(c *gin.Context) {
	var params deleteTaskReqParm
	if err := c.ShouldBindUri(&params); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	objectId, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	userOId := middleware.GetUserOId(c)

	if err := controller.DeleteTask(&objectId, userOId); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, gin.H{})
}

package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/middleware"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitTaskRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "", createTask)
	RouterPerms(routerGroup, http.MethodPost, "/find", findTask)
	RouterPerms(routerGroup, http.MethodPatch, "/:id/:checked", checkTask)
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

func checkTask(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	checkedStr := c.Param("checked")
	checked, err := strconv.ParseBool(checkedStr)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError("checked should be boolean"))
		return
	}

	userOId := middleware.GetUserOId(c)

	if err := controller.CheckTask(&objectId, userOId, &checked); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, gin.H{})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
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

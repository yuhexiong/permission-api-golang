package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitUserPermissionRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "", createUserPermission)
	RouterPerms(routerGroup, http.MethodPost, "/find", findUserPermission)
	RouterPerms(routerGroup, http.MethodDelete, "/:id", deleteUserPermission)
}

func createUserPermission(c *gin.Context) {
	var createUserPermissionOpts controller.CreateUserPermissionOpts

	if err := c.ShouldBindJSON(&createUserPermissionOpts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	userPermission := model.MapUserPermission{}
	err := controller.CreateUserPermission(createUserPermissionOpts, &userPermission)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, userPermission)
}

func findUserPermission(c *gin.Context) {
	var opts controller.FindUserPermissionOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	userPermissions := []*model.MapUserPermission{}
	if err := controller.FindUserPermission(opts, &userPermissions); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, userPermissions)
}

func deleteUserPermission(c *gin.Context) {
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	if err := controller.DeleteUserPermission(&objectId); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, gin.H{})
}

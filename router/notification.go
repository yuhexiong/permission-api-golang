package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/middleware"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitNotificationRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPatch, "/:id/read", readNotification)
}

type readNotificationReqParm struct {
	ID string `uri:"id" binding:"required"`
}

func readNotification(c *gin.Context) {
	var params readNotificationReqParm
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

	if err := controller.ReadNotification(&objectId, userOId); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, gin.H{})
}

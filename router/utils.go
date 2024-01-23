package router

import (
	"permission-api/middleware"

	"github.com/gin-gonic/gin"
)

func RouterPerms(router *gin.RouterGroup, method, path string, handler gin.HandlerFunc) {
	router.Handle(method, path, middleware.PermissionMiddle(), handler)
}

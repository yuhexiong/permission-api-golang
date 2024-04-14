package router

import (
	"permission-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(corsConfig()))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middleware.ErrorHandler())

	routerWithAuth := router.Group("")
	routerWithAuth.Use(middleware.AuthorizeToken)

	// auth 不需驗證 token
	InitUnAuthRouter(router.Group("/auth"))
	// auth 需驗證 token
	InitAuthRouter(routerWithAuth.Group("/auth"))

	// 其餘需驗證 token
	InitUserRouter(routerWithAuth.Group("/user"))
	InitPermissionRouter(routerWithAuth.Group("/permission"))
	InitUserPermissionRouter(routerWithAuth.Group("/userPermission"))
	InitTaskRouter(routerWithAuth.Group("/task"))
	InitNotificationRouter(routerWithAuth.Group("/notification"))
	InitSettingRouter(routerWithAuth.Group("/setting"))

	return router
}

func corsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowBrowserExtensions = true
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "Content-Length", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host"}

	return config
}

func RouterPerms(router *gin.RouterGroup, method, path string, handler gin.HandlerFunc) {
	router.Handle(method, path, middleware.PermissionMiddle(), handler)
}

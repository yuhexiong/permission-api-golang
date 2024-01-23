package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(corsConfig()))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	InitUserRouter(router.Group("/user"))
	InitPermissionRouter(router.Group("/permission"))

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

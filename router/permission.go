package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitPermissionRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/find", findPermission)
	RouterPerms(routerGroup, http.MethodPost, "", createPermission)
	// RouterPerms(routerGroup, http.MethodPatch, "/:id", updatePermission)
	RouterPerms(routerGroup, http.MethodPatch, "/enable/:id", enablePermission)
	RouterPerms(routerGroup, http.MethodDelete, "/:id", deletePermission)
}

func findPermission(c *gin.Context) {

}

func createPermission(c *gin.Context) {

}

func enablePermission(c *gin.Context) {

}

func deletePermission(c *gin.Context) {

}

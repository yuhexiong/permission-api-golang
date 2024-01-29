package router

import (
	"net/http"
	"permission-api/controller/permissionController"
	"permission-api/model"
	"permission-api/response"

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
	var opts permissionController.FindPermissionOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	permissions := []*model.Permission{}
	if err := permissionController.FindPermission(opts, &permissions); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response.ResFormat(c, http.StatusOK, 0, permissions)
}

func createPermission(c *gin.Context) {

}

func enablePermission(c *gin.Context) {

}

func deletePermission(c *gin.Context) {

}

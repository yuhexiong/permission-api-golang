package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodPost, "/login", login)
}

func login(c *gin.Context) {

}

package router

import (
	"net/http"
	"permission-api/controller"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"

	"github.com/gin-gonic/gin"
)

func InitSettingRouter(routerGroup *gin.RouterGroup) {
	RouterPerms(routerGroup, http.MethodGet, "/:code", getSettingByCode)
	RouterPerms(routerGroup, http.MethodPatch, "/:code/:value", updateSettingValue)
}

type getSettingByCodeReqParm struct {
	Code string `uri:"code" binding:"required"`
}

func getSettingByCode(c *gin.Context) {
	var params getSettingByCodeReqParm
	if err := c.ShouldBindUri(&params); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	setting := model.Setting{}
	if err := controller.GetSettingByCode(&params.Code, &setting); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, setting)
	c.Next()
}

type updateSettingValueReqParm struct {
	Code  string `uri:"code" binding:"required"`
	Value string `uri:"value" binding:"required"`
}

func updateSettingValue(c *gin.Context) {
	var params updateSettingValueReqParm
	if err := c.ShouldBindUri(&params); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	if err := controller.UpdateSettingValue(&params.Code, &params.Value); err != nil {
		response.AbortError(c, util.InvalidParameterError(err.Error()))
		return
	}

	response.SuccessFormat(c, nil)
	c.Next()
}

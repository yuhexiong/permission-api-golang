package response

import (
	"net/http"
	"permission-api/util"

	"github.com/gin-gonic/gin"
)

type JsonResult struct {
	Code int `json:"code"` // 錯誤代號，無異常時為 0
	Data any `json:"data"` // 結果資料
}

func SuccessFormat(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, JsonResult{
		Code: 0,
		Data: data,
	})
}

func AbortError(c *gin.Context, err error) {
	apiError, ok := err.(util.ErrorFormat)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "000001", "message": err.Error()})
		c.Abort()
	}

	c.JSON(apiError.StatusCode, gin.H{"code": apiError.Code, "message": apiError.Message})
	c.Abort()
}

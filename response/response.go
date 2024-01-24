package response

import "github.com/gin-gonic/gin"

type JsonResult struct {
	Code int `json:"code"` // 錯誤代號，無異常時為 0
	Data any `json:"data"` // 結果資料
}

// Error codes
const (
	BodyParserError  = "001000"
	InvalidParameter = "000108"
	InvalidDbData    = "000300"
)

// new error response
func ResFormat(c *gin.Context, statusCode int, code int, data interface{}) {
	c.JSON(statusCode, JsonResult{
		Code: code,
		Data: data,
	})
}

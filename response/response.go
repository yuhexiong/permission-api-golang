package response

import "github.com/gin-gonic/gin"

type UserResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error codes
const (
	BodyParserError  = "001000"
	InvalidParameter = "000108"
	InvalidDbData    = "000300"
)

// new error response
func NewErrorResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, UserResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

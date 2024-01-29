package middleware

import (
	"permission-api/util"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		lastErr := c.Errors.Last()
		if lastErr != nil {
			if apiError, ok := errors.Cause(lastErr.Err).(util.ErrorFormat); ok {
				c.JSON(apiError.StatusCode, gin.H{"code": apiError.Code, "message": apiError.Message})
				return
			}
		}

		c.Next()
	}
}

package middleware

import (
	"context"
	"net/http"
	"permission-api/util"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type contextKey string

const (
	errorCodeKey    contextKey = "errorCode"
	errorMessageKey contextKey = "errorMessage"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		statusCode := http.StatusInternalServerError
		errorCode := "000001"
		errorMessage := ""

		if lastErr := c.Errors.Last(); lastErr != nil {
			if apiError, ok := errors.Cause(lastErr.Err).(util.ErrorFormat); ok {
				errorCode = apiError.Code
				errorMessage = apiError.Message
			} else {
				errorMessage = lastErr.Err.Error()
			}

			c.JSON(statusCode, gin.H{"code": errorCode, "message": errorMessage})

			// save err for saveLogMiddleware
			ctx := context.WithValue(c.Request.Context(), errorCodeKey, errorCode)
			ctx = context.WithValue(ctx, errorMessageKey, errorMessage)
			c.Request = c.Request.WithContext(ctx)

		}

		c.Next()
	}
}

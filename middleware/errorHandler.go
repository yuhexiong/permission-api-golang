package middleware

import (
	"fmt"
	"permission-api/response"
	"permission-api/util"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {

			if err := recover(); err != nil {
				errorMessage := "Unknown error"
				if e, ok := err.(error); ok {
					errorMessage = fmt.Sprintf("%+v", e)
					stackTrace := fmt.Sprintf("%+v", errors.WithStack(e))
					util.RedLog("Stack Trace: %s\n", stackTrace)
				}

				response.AbortError(c, util.InternalServerError(errorMessage))
			}

		}()

		c.Next()
	}
}

package middleware

import (
	"permission-api/model"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := make(map[string]interface{})
		c.Bind(&requestBody)

		errorCodeInterface := c.Request.Context().Value(errorCodeKey)
		errorMessageInterface := c.Request.Context().Value(errorMessageKey)

		var errorCode, errorMessage *string
		if errorCodeInterface != nil {
			if code, ok := errorCodeInterface.(string); ok {
				errorCode = &code
			}
		}

		if errorMessageInterface != nil {
			if msg, ok := errorMessageInterface.(string); ok {
				errorMessage = &msg
			}
		}

		logData := model.Log{
			Method:       c.Request.Method,
			URL:          c.Request.URL.String(),
			Headers:      c.Request.Header,
			UserOID:      GetUserOId(c),
			StatusCode:   c.Writer.Status(),
			RequestBody:  requestBody,
			ErrorCode:    errorCode,
			ErrorMessage: errorMessage,
			CreatedAt:    time.Now(),
		}

		err := model.Insert(model.LogCollName, logData, nil)
		if err != nil {
			panic(err)
		}
	}
}

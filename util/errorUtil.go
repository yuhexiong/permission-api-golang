package util

import "net/http"

type ErrorFormat struct {
	StatusCode int    `json:"statusCode"` // http狀態碼
	Code       string `json:"code"`       // 錯誤代號，無異常時為 0
	Message    string `json:"message"`    // 錯誤訊息
}

func (apiError ErrorFormat) Error() string {
	return apiError.Message
}

func InvalidTokenError(message string) ErrorFormat {
	return ErrorFormat{StatusCode: http.StatusBadRequest, Code: "000002", Message: "[InvalidToken]" + message}
}

func UserNotFoundError(message string) ErrorFormat {
	return ErrorFormat{StatusCode: http.StatusBadRequest, Code: "000003", Message: "[UserNotFound]" + message}
}

func PermissionDeniedError(message string) ErrorFormat {
	return ErrorFormat{StatusCode: http.StatusBadRequest, Code: "000107", Message: "[PermissionDenied]" + message}
}

func InvalidParameterError(message string) ErrorFormat {
	return ErrorFormat{StatusCode: http.StatusBadRequest, Code: "000108", Message: "[InvalidParameter]" + message}
}

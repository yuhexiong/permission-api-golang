package util

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

// 驗證密碼
func ValidatePassword(userPasswordHash string, userPasswordSalt string, password string) bool {
	return userPasswordHash == HashPasswordWithSalt(password, userPasswordSalt)
}

// 加密密碼
func HashPasswordWithSalt(password string, salt string) string {
	h := hmac.New(sha512.New, []byte(salt))
	h.Write([]byte(password))

	return hex.EncodeToString(h.Sum(nil))
}

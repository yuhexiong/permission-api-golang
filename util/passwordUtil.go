package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"math"
	"time"
)

var SystemTokenLifeTime = 3153600000000000000 * time.Nanosecond
var NormalTokenLifeTime = 86400000000000 * time.Nanosecond

// 產生16進位字串
func GenerateHex(n int) (string, error) {
	b := make([]byte, int(math.Ceil(float64(n)/2)))
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// 加密密碼
func HashPasswordWithSalt(password string, salt string) string {
	return hex.EncodeToString(hmac.New(sha512.New, []byte(salt)).Sum([]byte(password)))
}

// 驗證密碼
func ValidatePassword(userPasswordHash string, userPasswordSalt string, password string) bool {
	return userPasswordHash == HashPasswordWithSalt(password, userPasswordSalt)
}

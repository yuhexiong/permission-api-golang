package middleware

import (
	"os"
	"permission-api/controller/permissionController"
	"permission-api/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey string

func init() {
	jwtKey = os.Getenv("JWTKey")
}

type Claims struct {
	jwt.StandardClaims
	UserId       *primitive.ObjectID
	PasswordHash string
}

// 產生加密過後的 token
func CreateToken(user *model.User, signedString *string) error {
	var expiresAt int64

	if model.UserTypeOpt(user.UserType) == model.UserTypeSystem {
		expiresAt = time.Now().Add(3153600000000000000).Unix() // 100年
	} else {
		expiresAt = time.Now().Add(86400000000000).Unix() // 1天
	}

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		UserId:       user.ID,
		PasswordHash: user.PasswordHash,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return err
	}

	*signedString = signedToken

	return nil
}

type PermissionInfo struct {
	UserOId       string                                          `json:"userOId" example:"623853b9503ce2ecdd221c94"` // 始俑者 ObjectId
	PermissionMap *map[string][]permissionController.PermissionOp `json:"-"`                                          // 權限key: [category]-[code] value: "R","W"
}

func SetPermissionInfo(c *gin.Context, permissionInfo *PermissionInfo) {
	if permissionInfo != nil {
		c.Set("permissionInfo", permissionInfo)
	}
}

func GetPermissionInfo(c *gin.Context) *PermissionInfo {
	permissionInfo, exists := c.Get("permissionInfo")

	if !exists {
		return nil
	}

	return permissionInfo.(*PermissionInfo)
}

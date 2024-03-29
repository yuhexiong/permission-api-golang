package middleware

import (
	"os"
	"permission-api/controller/permissionController"
	"permission-api/controller/sessionController"
	"permission-api/model"
	"permission-api/response"
	"permission-api/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey string

func init() {
	jwtKey = os.Getenv("JWTKey")
}

func AuthorizeToken(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	splits := strings.Split(auth, "Bearer ")
	if len(splits) < 2 {
		response.AbortError(c, util.InvalidParameterError("on token format"))
		return
	}

	token := splits[1]
	var session model.Session
	err := sessionController.GetSessionByToken(token, &session)
	if err != nil || session.SessionToken == "" {
		response.AbortError(c, util.InvalidParameterError("on get session by token"))
		return
	}

	permissionMap, err := permissionController.GetPermissionInfoByUser(session.UserOId)
	if err != nil {
		response.AbortError(c, util.InvalidParameterError("on get permissionMap by user"))
		return
	}

	setSessionToken(c, token)
	setUserOId(c, session.UserOId)
	SetPermissionMap(c, permissionMap)

	c.Next()
}

type Claims struct {
	jwt.StandardClaims
	UserId       *primitive.ObjectID
	PasswordHash string
}

// 產生加密過後的 token
func CreateToken(user *model.User, signedString *string) error {
	var expiresAt int64

	if model.UserType(user.UserType) == model.UserTypeSystem {
		expiresAt = time.Now().Add(util.SystemTokenLifeTime).Unix() // 100年
	} else {
		expiresAt = time.Now().Add(util.NormalTokenLifeTime).Unix() // 1天
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

func setUserOId(c *gin.Context, userOId *primitive.ObjectID) {
	c.Set("userOId", userOId)
}

func GetUserOId(c *gin.Context) *primitive.ObjectID {
	userOId, exists := c.Get("userOId")
	if !exists {
		return nil
	}

	return userOId.(*primitive.ObjectID)
}

func setSessionToken(c *gin.Context, token string) {
	c.Set("sessionToken", token)
}

func GetSessionToken(c *gin.Context) string {
	return c.GetString("sessionToken")
}

func SetPermissionMap(c *gin.Context, permissionMap *map[string][]model.PermissionOp) {
	if len(*permissionMap) > 0 {
		c.Set("permissionMap", permissionMap)
	}
}

func GetPermissionMap(c *gin.Context) *map[string][]model.PermissionOp {
	permissionMap, exists := c.Get("permissionMap")

	if !exists {
		return nil
	}

	return permissionMap.(*map[string][]model.PermissionOp)
}

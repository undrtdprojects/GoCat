package middlewares

import (
	"GoCat/helpers/common"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	RoleId   int    `json:"role_id"`
	jwt.StandardClaims
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			common.GenerateErrorResponse(c, err.Error())
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt_secret_key")), nil
		})
		if err != nil {
			common.GenerateErrorResponse(c, err.Error())
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("user", claims)
			c.Next()
		} else {
			common.GenerateErrorResponse(c, "Invalid token")
			return
		}
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if common.IsEmptyField(authHeader) {
		return tokenString, errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func GenerateJwtToken(userId int, username string, roleId int) (token string, err error) {
	// set token expiration time
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UserId:   userId,
		Username: username,
		RoleId:   roleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	GenerateJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = GenerateJwtToken.SignedString([]byte(viper.GetString("jwt_secret_key")))
	if err != nil {
		return
	}
	return
}

package middlewares

import (
	"GoCat/helpers/common"

	"github.com/gin-gonic/gin"
)

func RoleCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userCtx, exists := c.Get("user")
		if !exists {
			common.GenerateErrorResponse(c, "Unauthorized")
			c.Abort()
			return
		}

		user := userCtx.(*Claims)

		if user.RoleId == 2 {
			if c.Request.Method != "GET" {
				common.GenerateErrorResponse(c, "Read-only access")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

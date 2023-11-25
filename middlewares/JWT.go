package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"goserver/libs/e"
	"goserver/libs/jwt"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}

		code := e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			if claims, err := jwt.ParseToken(token); err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				if time.Now().Unix() > claims.ExpiresAt {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				} else {
					// 刷新token
					refreshToken := jwt.RefreshToken(*claims)
					c.Header("Authorization", refreshToken)
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": e.GetMsg(code),
				"data":    data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

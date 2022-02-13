package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	security "github.com/seishino/go-example/utils"
)

type JwtMiddleware struct {
	jwtUtil security.JWTAuthService
}

func NewJwtMiddleware(
	jwtUtil security.JWTAuthService,
) JwtMiddleware {
	return JwtMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (JwtMiddleware JwtMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 && t[0] == "Bearer" {
			authToken := t[1]
			authorized, err := JwtMiddleware.jwtUtil.Authorize(authToken)
			if authorized {
				c.Next()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not authorized",
		})
		c.Abort()
	}
}

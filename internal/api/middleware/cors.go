package middleware

import (
	"mifer/pkg/conf"

	"github.com/gin-gonic/gin"
)

// CorsMiddleware 跨域中间件
func CorsMiddleware(config *conf.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, origin := range config.Gin.Cors.AllowOrigins {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		for _, method := range config.Gin.Cors.AllowMethods {
			c.Writer.Header().Set("Access-Control-Allow-Methods", method)
		}
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

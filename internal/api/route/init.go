package route

import (
	"bloger/internal/api/handler"
	"bloger/internal/api/middleware"
	"bloger/internal/repo"
	"bloger/internal/service"
	"bloger/pkg/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route 路由
type Route struct {
	handler *handler.Handler
	config  *conf.Config
}

// NewRoute 新建路由
func (r *Route) NewRoute(config *conf.Config) error {
	repo := repo.NewRepo()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	r.handler = handler
	r.config = config
	return nil
}

// Setup 设置路由
func (r *Route) Setup() *gin.Engine {
	gin.SetMode(r.config.Gin.Mode)

	route := gin.Default()
	route.Use(middleware.CorsMiddleware(r.config))
	// 路由
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	return route
}

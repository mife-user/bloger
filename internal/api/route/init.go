package route

import (
	"bloger/internal/api/handler/githandler"
	"bloger/internal/api/middleware"
	"bloger/internal/repo/gitrepo"
	"bloger/internal/service/gitservice"
	"bloger/pkg/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route 路由
type Route struct {
	githandler *githandler.GitHandler
	config     *conf.Config
}

// NewRoute 新建路由
func (r *Route) NewRoute(config *conf.Config) error {
	repo := gitrepo.NewGitRepo(config)
	service := gitservice.NewGitService(repo)
	handler := githandler.NewGitHandler(service)
	r.githandler = handler
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
	// 保存token
	route.POST("/git/save", r.githandler.Save)
	return route
}

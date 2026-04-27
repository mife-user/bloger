package route

import (
	"context"

	"mifer/internal/ai/executor"
	"mifer/internal/api/handler/agenthandler"
	"mifer/internal/api/handler/githandler"
	"mifer/internal/api/middleware"
	"mifer/internal/repo/gitrepo"
	"mifer/internal/service/agentservice"
	"mifer/internal/service/gitservice"
	"mifer/pkg/conf"

	"github.com/gin-gonic/gin"
)

// Route 路由
type Route struct {
	githandler   *githandler.GitHandler
	agenthandler *agenthandler.AgentHandler
	config       *conf.Config
}

// NewRoute 新建路由
func (r *Route) NewRoute(config *conf.Config) error {

	repo := gitrepo.NewGitRepo(config)
	gitservice := gitservice.NewGitService(repo)
	githandler := githandler.NewGitHandler(gitservice)
	r.githandler = githandler

	executor, err := executor.InitExecutor(context.Background(), config)
	if err != nil {
		return err
	}
	executorService := agentservice.NewAgentService(executor)
	agentHandler := agenthandler.NewAgentHandler(executorService)
	r.agenthandler = agentHandler
	r.config = config
	return nil
}

// Setup 设置路由
func (r *Route) Setup() *gin.Engine {
	gin.SetMode(r.config.Gin.Mode)

	route := gin.Default()
	route.Use(middleware.CorsMiddleware(r.config))
	// 保存token
	gitRoute := route.Group("/git")
	{
		gitRoute.POST("/save", r.githandler.Save)
	}
	// 调用模型
	agentRoute := route.Group("/agent")
	{
		agentRoute.POST("/chat", r.agenthandler.Chat)
	}
	return route
}

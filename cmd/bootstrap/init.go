package bootstrap

import (
	"bloger/internal/api/route"
	"bloger/pkg/conf"
	"bloger/pkg/logger"
	"context"

	"github.com/gin-gonic/gin"
)

// App 应用
type App struct {
	engine *gin.Engine
	config *conf.Config
	route  *route.Route
}

// Init 初始化应用
func Init() (*App, error) {
	var app App
	// 加载配置
	if err := app.LoadConfig(); err != nil {
		return nil, err
	}
	// 加载日志
	if err := app.LoadLogger(); err != nil {
		logger.Error("LoadLogger failed", logger.C(err))
		return nil, err
	}
	// 加载路由
	if err := app.LoadRoute(); err != nil {
		logger.Error("LoadRoute failed", logger.C(err))
		return nil, err
	}

	return &app, nil
}

// Run 运行应用
func (a *App) Run() error {
	logger.Info("Run app:", logger.S("port", a.config.Gin.Port))
	return a.engine.Run(a.config.Gin.Port)
}

func (a *App) Down(ctx context.Context) error {
	logger.Info("应用正在关闭...")
	return nil
}

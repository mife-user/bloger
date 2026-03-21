package bootstrap

import (
	"bloger/pkg/conf"
	"bloger/pkg/logger"
)

// LoadRoute 加载路由
func (a *App) LoadRoute() error {
	if err := a.route.NewRoute(a.config); err != nil {
		logger.Error("LoadRoute failed", logger.C(err))
		return err
	}
	a.engine = a.route.Setup()
	return nil
}

// LoadConfig 加载配置
func (a *App) LoadConfig() error {
	if err := conf.LoadConfig(); err != nil {
		logger.Error("LoadConfig failed", logger.C(err))
		return err
	}
	a.config = conf.GetConfig()
	return nil
}

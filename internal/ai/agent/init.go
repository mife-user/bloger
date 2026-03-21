package agent

import (
	"bloger/pkg/logger"

	"github.com/tmc/langchaingo/agents"
)

type Agent struct {
	agent *agents.Agent
}

func InitAgent() {
	logger.Info("初始化智能体...")
}

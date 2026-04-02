package agenter

import (
	"bloger/pkg/logger"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/tools"
)

// InitAgent 初始化Agent
func InitAgent(llm llms.Model, agentTools []tools.Tool) agents.Agent {
	logger.Info("初始化Agent...")
	agent := agents.NewOpenAIFunctionsAgent(llm, agentTools)
	return agent
}

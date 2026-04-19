package agenter

import (
	"bloger/pkg/logger"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

// InitAgent 初始化Agent
func InitAgent(llm llms.Model, prompter *prompts.PromptTemplate, agentTools []tools.Tool) agents.Agent {
	logger.Info("初始化Agent...")

	// 创建 ConversationalAgent
	agent := agents.NewConversationalAgent(llm, agentTools, agents.WithPrompt(*prompter))

	return agent
}

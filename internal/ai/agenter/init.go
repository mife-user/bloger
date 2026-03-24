package agenter

import (
	"bloger/pkg/logger"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

// InitAgent 初始化Agent
func InitAgent(llm llms.Model, tools []tools.Tool, agentPrompt prompts.PromptTemplate) *agents.ConversationalAgent {
	logger.Info("初始化Agent...")
	agent := agents.NewConversationalAgent(llm, tools, agents.WithPrompt(agentPrompt))
	return agent
}

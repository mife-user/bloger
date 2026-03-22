package agent

import (
	"bloger/internal/ai/llm"
	"bloger/pkg/logger"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/memory"
)

type Agent struct {
	agent *agents.ConversationalAgent
}

// InitAgent 初始化Agent
func InitAgent(llm *llm.LLM) *Agent {
	logger.Info("初始化Agent...")
	mem := agents.WithMemory(memory.NewConversationTokenBuffer(llm.Client, 60000))
	agent := agents.NewConversationalAgent(llm.Client, nil, mem)
	return &Agent{agent: agent}
}

package executor

import (
	"bloger/internal/ai/agenter"
	"bloger/internal/ai/llmer"
	"bloger/internal/ai/memoryer"
	"bloger/internal/ai/prompter"
	"bloger/internal/ai/tooler"
	"bloger/pkg/conf"

	"github.com/tmc/langchaingo/agents"
)

type Executor struct {
	executor *agents.Executor
}

func NewExecutor(executor *agents.Executor) *Executor {
	return &Executor{executor: executor}
}

// InitExecutor 初始化执行器
func InitExecutor(config *conf.Config) (*Executor, error) {
	// 初始化LLM
	llm, err := llmer.InitLLM(config)
	if err != nil {
		return nil, err
	}
	// 初始化工具
	agentTools := tooler.InitTools()
	// 初始化内存
	memory := memoryer.InitMemoryer(llm)
	// 初始化提示词
	prompter := prompter.InitPrompter(llm, config.Ai.SystemPrompt)
	// 初始化Agent
	agent := agenter.InitAgent(llm, &prompter, agentTools)

	// 创建Executor并设置内存
	executor := agents.NewExecutor(agent, agents.WithMemory(memory))
	return NewExecutor(executor), nil
}

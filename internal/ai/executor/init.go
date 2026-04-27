package executor

import (
	"context"

	"mifer/internal/ai/agenter"
	"mifer/internal/ai/llmer"
	"mifer/internal/ai/memoryer"
	"mifer/internal/ai/prompter"
	"mifer/internal/ai/tooler"
	"mifer/internal/domain"
	"mifer/pkg/conf"
	"mifer/pkg/logger"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

type Executor struct {
	agent       *react.Agent
	chatHistory *memoryer.ChatHistory
}

func InitExecutor(ctx context.Context, config *conf.Config) (*Executor, error) {
	llm, err := llmer.InitLLM(ctx, config)
	if err != nil {
		return nil, err
	}
	tools := tooler.InitTools()
	modifier := prompter.ModifierBuilder{}.Build(config.Ai.SystemPrompt)
	agent, err := agenter.InitAgent(ctx, llm, tools, modifier)
	if err != nil {
		return nil, err
	}
	return &Executor{
		agent:       agent,
		chatHistory: memoryer.New(2048),
	}, nil
}

func (e *Executor) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	e.chatHistory.Append(schema.UserMessage(input.Content))
	msg, err := e.agent.Generate(ctx, e.chatHistory.Messages())
	if err != nil {
		logger.Error("agent.Generate failed", logger.C(err))
		return domain.ChatResponse{}, err
	}
	e.chatHistory.Append(msg)
	return domain.ChatResponse{Content: msg.Content}, nil
}

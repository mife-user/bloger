package agentservice

import (
	"bloger/internal/domain"
	"bloger/pkg/task"
	"context"
)

// Chat 聊天
func (a *AgentService) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	var ChatResponse domain.ChatResponse
	err := task.Do(ctx,func() error{
		var err error
		ChatResponse,err =  a.agent.Chat(ctx, input)
		return err
	}
	)
}

package agentservice

import (
	"bloger/internal/domain"
	"context"
)

// Chat 聊天
func (a *AgentService) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	return a.agent.Chat(ctx, input)
}

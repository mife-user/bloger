package agentservice

import (
	"context"
)

// Chat 聊天
func (a *AgentService) Chat(ctx context.Context, input map[string]any) (map[string]any, error) {
	return a.agent.Chat(ctx, input)
}

package agent

import (
	"context"
)

// Chat 聊天
func (a *Agent) Chat(ctx context.Context, input map[string]any) (map[string]any, error) {
	// 调用链
	return a.agent.Chain.Call(ctx, input)
}

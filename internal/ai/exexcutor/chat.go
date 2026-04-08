package exexcutor

import (
	"context"

	"github.com/tmc/langchaingo/chains"
)

func (e *Executor) Chat(ctx context.Context, input map[string]any) (map[string]any, error) {
	return chains.Call(ctx, e.executor, input)
}

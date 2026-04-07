package exexcutor

import (
	"context"

	"github.com/tmc/langchaingo/chains"
)

func (e *Executor) Chat(ctx context.Context, input map[string]any) (string, error) {
	return chains.Run(ctx, e.executor, input)
}

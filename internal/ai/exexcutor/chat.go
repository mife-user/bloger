package exexcutor

import (
	"context"
)

func (e *Executor) Chat(ctx context.Context, input map[string]any) (map[string]any, error) {
	return e.executor.Call(ctx, input)
}

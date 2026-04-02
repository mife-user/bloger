package exexcutor

import (
	"context"
)

func (e *Executor) Chat(ctx context.Context, input map[string]any) (map[string]any, error) {
	// 调用执行器
	response, err := e.executor.Call(ctx, input)
	if err != nil {
		return nil, err
	}
	return response, nil
}

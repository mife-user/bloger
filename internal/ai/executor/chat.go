package executor

import (
	"bloger/internal/domain"
	"bloger/pkg/logger"
	"context"

	"github.com/tmc/langchaingo/chains"
)

func (e *Executor) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	var err error
	response, err := chains.Call(ctx, e.executor, input.Message)
	if err != nil {
		logger.Error("Call failed", logger.C(err))
		return domain.ChatResponse{}, err
	}
	err = e.executor.Memory.SaveContext(ctx, input.Message, response)
	if err != nil {
		logger.Error("SaveContext failed", logger.C(err))
		return domain.ChatResponse{}, err
	}
	return domain.ChatResponse{
		Message: response,
	}, nil
}

package executor

import (
	"bloger/internal/domain"
	"context"

	"bloger/internal/model/agentmodel"

	"github.com/tmc/langchaingo/chains"
)

func (e *Executor) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	var err error
	var responseModel agentmodel.ChatModel
	responseModel.Message, err = chains.Call(ctx, e.executor, input.Message)
	return domain.ChatResponse{
		Message: responseModel.Message,
	}, err
}

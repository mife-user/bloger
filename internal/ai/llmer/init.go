package llmer

import (
	"context"

	"mifer/pkg/conf"
	"mifer/pkg/errs"
	"mifer/pkg/logger"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino-ext/components/model/openai"
)

func InitLLM(ctx context.Context, config *conf.Config) (model.ToolCallingChatModel, error) {
	logger.Info("初始化LLM...")
	if config.Ai.ApiKey == "" {
		return nil, errs.New("api_key is empty")
	}
	if config.Ai.BaseURL == "" {
		return nil, errs.New("base_url is empty")
	}
	return openai.NewChatModel(ctx, &openai.ChatModelConfig{
		Model:   config.Ai.Model,
		APIKey:  config.Ai.ApiKey,
		BaseURL: config.Ai.BaseURL,
	})
}

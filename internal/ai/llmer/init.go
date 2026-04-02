package llmer

import (
	"bloger/pkg/conf"
	"bloger/pkg/errs"
	"bloger/pkg/logger"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// InitLLM 初始化LLM客户端
func InitLLM(config *conf.Config) (llms.Model, error) {
	logger.Info("初始化LLM...")
	// 验证配置
	if config.Ai.ApiKey == "" {
		return nil, errs.New("api_key is empty")
	}
	// 设置DeepSeek的API地址
	if config.Ai.BaseURL == "" {
		return nil, errs.New("base_url is empty")
	}
	// 创建 langchaingo 的 OpenAI 客户端
	client, err := openai.New(
		openai.WithToken(config.Ai.ApiKey),
		openai.WithModel(config.Ai.Model),
		openai.WithBaseURL(config.Ai.BaseURL), // DeepSeek 兼容 OpenAI API
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

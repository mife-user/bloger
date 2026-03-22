package llm

import (
	"bloger/pkg/conf"
	"bloger/pkg/err"
	"bloger/pkg/logger"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// LLM 大语言模型客户端
type LLM struct {
	Client llms.Model
}

// InitLLM 初始化LLM客户端
func InitLLM(config *conf.Config) (*LLM, error) {
	logger.Info("初始化LLM...")
	// 验证配置
	if config.Ai.ApiKey == "" {
		return nil, err.New("api_key is empty")
	}
	// 设置DeepSeek的API地址
	if config.Ai.BaseURL == "" {
		return nil, err.New("base_url is empty")
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
	return &LLM{
		Client: client,
	}, nil
}

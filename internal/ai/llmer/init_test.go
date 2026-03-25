package llmer

import (
	"bloger/pkg/conf"
	"testing"
)

// TestInitLLM_Success 测试成功初始化LLM
func TestInitLLM_Success(t *testing.T) {
	config := &conf.Config{
		Ai: conf.AiConfig{
			ApiKey:       "test-api-key",
			BaseURL:      "https://api.deepseek.com/v1",
			Model:        "deepseek-chat",
			SystemPrompt: "You are a helpful assistant.",
		},
	}

	llm, err := InitLLM(config)
	if err != nil {
		t.Fatalf("初始化LLM失败: %v", err)
	}

	if llm == nil {
		t.Fatal("LLM不应该为nil")
	}
}

// TestInitLLM_EmptyAPIKey 测试空API Key
func TestInitLLM_EmptyAPIKey(t *testing.T) {
	config := &conf.Config{
		Ai: conf.AiConfig{
			ApiKey:       "",
			BaseURL:      "https://api.deepseek.com/v1",
			Model:        "deepseek-chat",
			SystemPrompt: "You are a helpful assistant.",
		},
	}

	llm, err := InitLLM(config)

	// 应该返回错误
	if err == nil {
		t.Fatal("空API Key应该返回错误")
	}

	// LLM应该为nil
	if llm != nil {
		t.Fatal("错误时LLM应该为nil")
	}
}

// TestInitLLM_EmptyBaseURL 测试空BaseURL
func TestInitLLM_EmptyBaseURL(t *testing.T) {
	config := &conf.Config{
		Ai: conf.AiConfig{
			ApiKey:       "test-api-key",
			BaseURL:      "",
			Model:        "deepseek-chat",
			SystemPrompt: "You are a helpful assistant.",
		},
	}

	llm, err := InitLLM(config)

	// 应该返回错误
	if err == nil {
		t.Fatal("空BaseURL应该返回错误")
	}

	// LLM应该为nil
	if llm != nil {
		t.Fatal("错误时LLM应该为nil")
	}
}

// TestInitLLM_DifferentModels 测试不同模型
func TestInitLLM_DifferentModels(t *testing.T) {
	models := []string{
		"deepseek-chat",
		"deepseek-coder",
		"gpt-4",
		"gpt-3.5-turbo",
	}

	for _, model := range models {
		config := &conf.Config{
			Ai: conf.AiConfig{
				ApiKey:       "test-api-key",
				BaseURL:      "https://api.deepseek.com/v1",
				Model:        model,
				SystemPrompt: "You are a helpful assistant.",
			},
		}

		llm, err := InitLLM(config)
		if err != nil {
			t.Errorf("模型 %s 初始化失败: %v", model, err)
			continue
		}

		if llm == nil {
			t.Errorf("模型 %s 的LLM不应该为nil", model)
		}
	}
}

// TestInitLLM_DifferentBaseURLs 测试不同的BaseURL
func TestInitLLM_DifferentBaseURLs(t *testing.T) {
	baseURLs := []string{
		"https://api.deepseek.com/v1",
		"https://api.openai.com/v1",
		"http://localhost:8000/v1",
		"https://custom-llm.example.com/api",
	}

	for _, baseURL := range baseURLs {
		config := &conf.Config{
			Ai: conf.AiConfig{
				ApiKey:       "test-api-key",
				BaseURL:      baseURL,
				Model:        "test-model",
				SystemPrompt: "You are a helpful assistant.",
			},
		}

		llm, err := InitLLM(config)
		if err != nil {
			t.Errorf("BaseURL %s 初始化失败: %v", baseURL, err)
			continue
		}

		if llm == nil {
			t.Errorf("BaseURL %s 的LLM不应该为nil", baseURL)
		}
	}
}

// TestInitLLM_NilConfig 测试nil配置
func TestInitLLM_NilConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// 预期会panic，因为nil配置会导致空指针
			t.Log("nil配置导致panic（预期行为）")
		}
	}()

	llm, err := InitLLM(nil)

	// 如果没有panic，应该返回错误或nil LLM
	if err == nil && llm != nil {
		t.Fatal("nil配置应该导致错误或nil LLM")
	}
}

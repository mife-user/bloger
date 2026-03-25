package exexcutor

import (
	"bloger/pkg/conf"
	"testing"
)

// TestInitExecutor 测试初始化执行器
func TestInitExecutor(t *testing.T) {
	config := &conf.Config{
		Ai: conf.AiConfig{
			ApiKey:       "test-api-key",
			BaseURL:      "https://api.deepseek.com/v1",
			Model:        "deepseek-chat",
			SystemPrompt: "You are a helpful assistant.",
		},
	}

	executor, err := InitExecutor(config)

	if err != nil {
		t.Fatalf("初始化执行器失败: %v", err)
	}

	if executor == nil {
		t.Fatal("执行器不应该为nil")
	}
}

// TestInitExecutor_EmptyAPIKey 测试空API Key
func TestInitExecutor_EmptyAPIKey(t *testing.T) {
	config := &conf.Config{
		Ai: conf.AiConfig{
			ApiKey:       "",
			BaseURL:      "https://api.deepseek.com/v1",
			Model:        "deepseek-chat",
			SystemPrompt: "You are a helpful assistant.",
		},
	}

	executor, err := InitExecutor(config)

	// 应该返回错误
	if err == nil {
		t.Fatal("空API Key应该返回错误")
	}

	// 执行器应该为nil
	if executor != nil {
		t.Fatal("错误时执行器应该为nil")
	}
}

// TestInitExecutor_EmptyBaseURL 测试空BaseURL
func TestInitExecutor_EmptyBaseURL(t *testing.T) {
	config := &conf.Config{
		Ai: conf.AiConfig{
			ApiKey:       "test-api-key",
			BaseURL:      "",
			Model:        "deepseek-chat",
			SystemPrompt: "You are a helpful assistant.",
		},
	}

	executor, err := InitExecutor(config)

	// 应该返回错误
	if err == nil {
		t.Fatal("空BaseURL应该返回错误")
	}

	// 执行器应该为nil
	if executor != nil {
		t.Fatal("错误时执行器应该为nil")
	}
}

// TestInitExecutor_NilConfig 测试nil配置
func TestInitExecutor_NilConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// 预期会panic，因为nil配置会导致空指针
			t.Log("nil配置导致panic（预期行为）")
		}
	}()

	executor, err := InitExecutor(nil)

	// 如果没有panic，应该返回错误或nil执行器
	if err == nil && executor != nil {
		t.Fatal("nil配置应该导致错误或nil执行器")
	}
}

// TestNewExecutor 测试创建执行器
func TestNewExecutor(t *testing.T) {
	// 注意：这里需要实际的agents.Executor
	// 由于创建需要完整的初始化流程，我们只测试NewExecutor函数
	// 实际的Executor创建在InitExecutor中测试

	// 这里我们测试NewExecutor是否能正确包装
	executor := NewExecutor(nil)

	// NewExecutor应该能处理nil
	if executor == nil {
		t.Fatal("NewExecutor不应该返回nil")
	}
}

// TestInitExecutor_CompleteConfig 测试完整配置
func TestInitExecutor_CompleteConfig(t *testing.T) {
	config := &conf.Config{
		Ai: conf.AiConfig{
			ApiKey:       "test-api-key-12345",
			BaseURL:      "https://api.deepseek.com/v1",
			Model:        "deepseek-chat",
			SystemPrompt: "You are an AI assistant for blog writing. Help users create engaging content.",
		},
	}

	executor, err := InitExecutor(config)

	if err != nil {
		t.Fatalf("初始化执行器失败: %v", err)
	}

	if executor == nil {
		t.Fatal("执行器不应该为nil")
	}

	// 验证executor的内部结构
	if executor.executor == nil {
		t.Fatal("内部executor不应该为nil")
	}
}

// TestInitExecutor_DifferentModels 测试不同模型
func TestInitExecutor_DifferentModels(t *testing.T) {
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

		executor, err := InitExecutor(config)

		if err != nil {
			t.Errorf("模型 %s 初始化失败: %v", model, err)
			continue
		}

		if executor == nil {
			t.Errorf("模型 %s 的执行器不应该为nil", model)
		}
	}
}

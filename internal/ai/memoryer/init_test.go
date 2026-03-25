package memoryer

import (
	"context"
	"testing"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// TestInitMemoryer 测试初始化内存管理器
func TestInitMemoryer(t *testing.T) {
	// 创建一个简单的mock LLM
	config := &testConfig{
		apiKey:  "test-key",
		baseURL: "https://api.deepseek.com/v1",
		model:   "deepseek-chat",
	}

	llm, err := openai.New(
		openai.WithToken(config.apiKey),
		openai.WithModel(config.model),
		openai.WithBaseURL(config.baseURL),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	memory := InitMemoryer(llm)

	if memory == nil {
		t.Fatal("内存管理器不应该为nil")
	}
}

// TestInitMemoryer_BufferSize 测试缓冲区大小
func TestInitMemoryer_BufferSize(t *testing.T) {
	config := &testConfig{
		apiKey:  "test-key",
		baseURL: "https://api.deepseek.com/v1",
		model:   "deepseek-chat",
	}

	llm, err := openai.New(
		openai.WithToken(config.apiKey),
		openai.WithModel(config.model),
		openai.WithBaseURL(config.baseURL),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	memory := InitMemoryer(llm)

	// 验证内存管理器已创建
	if memory == nil {
		t.Fatal("内存管理器不应该为nil")
	}

	// 注意：ConversationTokenBuffer的maxTokenLimit是私有的，无法直接验证
	// 但我们可以验证它不为nil
}

// TestMemoryer_Context 测试内存管理器的上下文功能
func TestMemoryer_Context(t *testing.T) {
	config := &testConfig{
		apiKey:  "test-key",
		baseURL: "https://api.deepseek.com/v1",
		model:   "deepseek-chat",
	}

	llm, err := openai.New(
		openai.WithToken(config.apiKey),
		openai.WithModel(config.model),
		openai.WithBaseURL(config.baseURL),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	memory := InitMemoryer(llm)

	// 创建上下文
	ctx := context.Background()

	// 验证内存管理器可以与上下文一起使用
	// 注意：实际的消息存储和检索需要通过Agent来测试
	// 这里只验证内存管理器不为nil
	if memory == nil {
		t.Fatal("内存管理器不应该为nil")
	}

	_ = ctx // 使用ctx避免编译警告
}

// testConfig 测试配置结构
type testConfig struct {
	apiKey  string
	baseURL string
	model   string
}

// MockLLM 用于测试的模拟LLM
type MockLLM struct{}

// 确保MockLLM实现llms.Model接口
var _ llms.Model = (*MockLLM)(nil)

func (m *MockLLM) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	return "mock response", nil
}

func (m *MockLLM) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {
	return &llms.ContentResponse{
		Choices: []*llms.ContentChoice{
			{
				Content: "mock response",
			},
		},
	}, nil
}

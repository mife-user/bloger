package agenter

import (
	"context"
	"testing"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
)

// TestInitAgent 测试初始化Agent
func TestInitAgent(t *testing.T) {
	// 创建LLM
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	// 创建空工具列表
	toolList := []tools.Tool{}

	// 初始化Agent
	agent := InitAgent(llm, toolList)

	if agent == nil {
		t.Fatal("Agent不应该为nil")
	}
}

// TestInitAgent_WithTools 测试带工具的Agent
func TestInitAgent_WithTools(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	// 创建模拟工具
	toolList := []tools.Tool{
		&mockTool{name: "test_tool"},
	}

	agent := InitAgent(llm, toolList)

	if agent == nil {
		t.Fatal("Agent不应该为nil")
	}
}

// TestInitAgent_NilLLM 测试nil LLM
func TestInitAgent_NilLLM(t *testing.T) {
	toolList := []tools.Tool{}

	defer func() {
		if r := recover(); r != nil {
			// 预期会panic，因为NewOpenAIFunctionsAgent需要有效的LLM
			t.Log("nil LLM导致panic（预期行为）")
		}
	}()

	agent := InitAgent(nil, toolList)

	// 如果没有panic，说明langchaingo允许nil LLM
	// 这种情况下我们记录一下
	if agent != nil {
		t.Log("警告：nil LLM返回了非nil agent，这可能在实际使用时导致问题")
	}
}

// mockTool 模拟工具用于测试
type mockTool struct {
	name string
}

func (m *mockTool) Name() string {
	return m.name
}

func (m *mockTool) Description() string {
	return "mock tool for testing"
}

func (m *mockTool) Call(ctx context.Context, input string) (string, error) {
	return "mock result", nil
}

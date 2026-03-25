package agenter

import (
	"context"
	"testing"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
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

	// 创建提示词模板
	prompt := prompts.NewPromptTemplate("You are a helpful assistant.", []string{"input"})

	// 初始化Agent
	agent := InitAgent(llm, toolList, prompt)

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

	prompt := prompts.NewPromptTemplate("You are a helpful assistant.", []string{"input"})

	agent := InitAgent(llm, toolList, prompt)

	if agent == nil {
		t.Fatal("Agent不应该为nil")
	}
}

// TestInitAgent_NilLLM 测试nil LLM
func TestInitAgent_NilLLM(t *testing.T) {
	toolList := []tools.Tool{}
	prompt := prompts.NewPromptTemplate("You are a helpful assistant.", []string{"input"})

	defer func() {
		if r := recover(); r != nil {
			// 预期会panic，因为NewConversationalAgent需要有效的LLM
			t.Log("nil LLM导致panic（预期行为）")
		}
	}()

	agent := InitAgent(nil, toolList, prompt)

	// 如果没有panic，说明langchaingo允许nil LLM
	// 这种情况下我们记录一下
	if agent != nil {
		t.Log("警告：nil LLM返回了非nil agent，这可能在实际使用时导致问题")
	}
}

// TestInitAgent_EmptyPrompt 测试空提示词
func TestInitAgent_EmptyPrompt(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	toolList := []tools.Tool{}

	// 空提示词
	prompt := prompts.NewPromptTemplate("", []string{"input"})

	agent := InitAgent(llm, toolList, prompt)

	// 空提示词也应该能创建Agent
	if agent == nil {
		t.Fatal("空提示词也应该能创建Agent")
	}
}

// TestInitAgent_ComplexPrompt 测试复杂提示词
func TestInitAgent_ComplexPrompt(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	toolList := []tools.Tool{}

	// 复杂提示词
	complexPrompt := `You are an AI assistant specialized in blog writing.

Your responsibilities:
1. Generate blog post ideas
2. Write engaging content
3. Optimize for SEO
4. Format in Markdown

Always be helpful and creative.`

	prompt := prompts.NewPromptTemplate(complexPrompt, []string{"input"})

	agent := InitAgent(llm, toolList, prompt)

	if agent == nil {
		t.Fatal("复杂提示词应该能创建Agent")
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

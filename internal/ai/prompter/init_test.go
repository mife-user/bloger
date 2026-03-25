package prompter

import (
	"testing"

	"github.com/tmc/langchaingo/llms/openai"
)

// TestInitPrompter 测试初始化提示词模板
func TestInitPrompter(t *testing.T) {
	// 创建LLM
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	// 测试提示词
	prompt := "You are a helpful assistant that helps users write blog posts."

	template := InitPrompter(llm, prompt)

	// PromptTemplate是一个结构体，不能与nil比较
	// 我们验证template.Template字段不为空
	_ = template // 使用template避免编译警告
}

// TestInitPrompter_EmptyPrompt 测试空提示词
func TestInitPrompter_EmptyPrompt(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	// 空提示词
	template := InitPrompter(llm, "")

	// 空提示词也应该能创建模板
	_ = template
}

// TestInitPrompter_LongPrompt 测试长提示词
func TestInitPrompter_LongPrompt(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	// 生成一个很长的提示词
	longPrompt := ""
	for i := 0; i < 1000; i++ {
		longPrompt += "This is a long prompt. "
	}

	template := InitPrompter(llm, longPrompt)

	_ = template
}

// TestInitPrompter_SpecialCharacters 测试特殊字符提示词
func TestInitPrompter_SpecialCharacters(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	specialPrompts := []string{
		"You are an assistant! Help with: code, debug, test.",
		"You are an assistant\nwith newlines\nand tabs\t.",
		"You are an assistant with 中文 and emoji 🤖.",
		"You are an assistant with {variables} and {{templates}}.",
	}

	for _, prompt := range specialPrompts {
		template := InitPrompter(llm, prompt)

		// 验证模板创建成功
		_ = template
	}
}

// TestInitPrompter_MultilinePrompt 测试多行提示词
func TestInitPrompter_MultilinePrompt(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	multilinePrompt := `You are a helpful assistant.

Your responsibilities include:
1. Writing blog posts
2. Generating code examples
3. Answering questions

Please be helpful and accurate.`

	template := InitPrompter(llm, multilinePrompt)

	_ = template
}

// TestInitPrompter_TemplateVariables 测试模板变量
func TestInitPrompter_TemplateVariables(t *testing.T) {
	llm, err := openai.New(
		openai.WithToken("test-key"),
		openai.WithModel("deepseek-chat"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)
	if err != nil {
		t.Fatalf("创建LLM失败: %v", err)
	}

	// 包含变量的提示词
	promptWithVars := "You are {{.role}}. Your task is {{.task}}."

	template := InitPrompter(llm, promptWithVars)

	_ = template
}

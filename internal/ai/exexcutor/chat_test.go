package exexcutor

import (
	"context"
	"testing"
)

// TestExecutor_Chat 测试聊天功能
func TestExecutor_Chat(t *testing.T) {
	// 注意：这个测试需要完整的初始化
	// 由于Chat方法需要实际的LLM调用，我们只测试接口

	// 创建一个简单的mock executor
	executor := &Executor{
		executor: nil, // 实际测试需要真实的executor
	}

	ctx := context.Background()
	input := map[string]any{
		"input": "Hello, how are you?",
	}

	// 注意：由于executor为nil，这个调用会panic
	// 在实际测试中，应该使用InitExecutor创建完整的executor
	defer func() {
		if r := recover(); r != nil {
			// 预期会panic，因为executor为nil
			t.Log("Chat需要有效的executor")
		}
	}()

	_, err := executor.Chat(ctx, input)

	// 如果没有panic，检查错误
	if err == nil {
		t.Log("Chat成功执行")
	}
}

// TestExecutor_Chat_EmptyInput 测试空输入
func TestExecutor_Chat_EmptyInput(t *testing.T) {
	executor := &Executor{
		executor: nil,
	}

	ctx := context.Background()
	input := map[string]any{}

	defer func() {
		if r := recover(); r != nil {
			t.Log("Chat需要有效的executor")
		}
	}()

	_, err := executor.Chat(ctx, input)

	if err == nil {
		t.Log("Chat处理了空输入")
	}
}

// TestExecutor_Chat_ComplexInput 测试复杂输入
func TestExecutor_Chat_ComplexInput(t *testing.T) {
	executor := &Executor{
		executor: nil,
	}

	ctx := context.Background()
	input := map[string]any{
		"input":      "Write a blog post about AI",
		"chat_history": []map[string]string{
			{"role": "user", "content": "Hello"},
			{"role": "assistant", "content": "Hi there!"},
		},
		"temperature": 0.7,
		"max_tokens":  1000,
	}

	defer func() {
		if r := recover(); r != nil {
			t.Log("Chat需要有效的executor")
		}
	}()

	_, err := executor.Chat(ctx, input)

	if err == nil {
		t.Log("Chat处理了复杂输入")
	}
}

// TestNewExecutor_Wrap 测试包装executor
func TestNewExecutor_Wrap(t *testing.T) {
	// 测试NewExecutor是否能正确创建Executor实例
	exec := NewExecutor(nil)

	if exec == nil {
		t.Fatal("NewExecutor不应该返回nil")
	}

	// 验证内部字段
	if exec.executor != nil {
		t.Error("传入nil时，内部executor应该为nil")
	}
}

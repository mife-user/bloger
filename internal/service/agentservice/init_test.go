package agentservice

import (
	"context"
	"testing"
)

// TestNewAgentService 测试创建AgentService
func TestNewAgentService(t *testing.T) {
	// 创建mock agent
	mockAgent := &mockAgent{}

	service := NewAgentService(mockAgent)

	if service == nil {
		t.Fatal("AgentService不应该为nil")
	}

	if service.agent == nil {
		t.Fatal("AgentService.agent不应该为nil")
	}
}

// TestAgentService_Chat 测试聊天功能
func TestAgentService_Chat(t *testing.T) {
	mockAgent := &mockAgent{
		response: map[string]any{"output": "Hello! How can I help you?"},
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()
	input := map[string]any{
		"input": "Hello",
	}

	result, err := service.Chat(ctx, input)

	if err != nil {
		t.Fatalf("Chat失败: %v", err)
	}

	if result == nil {
		t.Fatal("结果不应该为nil")
	}

	// 验证响应
	if output, ok := result["output"]; ok {
		if output != "Hello! How can I help you?" {
			t.Errorf("期望输出 'Hello! How can I help you?', 得到 %v", output)
		}
	}
}

// TestAgentService_Chat_Error 测试聊天错误处理
func TestAgentService_Chat_Error(t *testing.T) {
	mockAgent := &mockAgent{
		err: context.DeadlineExceeded,
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()
	input := map[string]any{
		"input": "Hello",
	}

	_, err := service.Chat(ctx, input)

	if err == nil {
		t.Fatal("应该返回错误")
	}

	if err != context.DeadlineExceeded {
		t.Errorf("期望错误 '%v', 得到 '%v'", context.DeadlineExceeded, err)
	}
}

// TestAgentService_Chat_EmptyInput 测试空输入
func TestAgentService_Chat_EmptyInput(t *testing.T) {
	mockAgent := &mockAgent{
		response: map[string]any{"output": "Please provide input"},
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()
	input := map[string]any{}

	result, err := service.Chat(ctx, input)

	if err != nil {
		t.Fatalf("Chat失败: %v", err)
	}

	if result == nil {
		t.Fatal("结果不应该为nil")
	}
}

// TestAgentService_Chat_ComplexInput 测试复杂输入
func TestAgentService_Chat_ComplexInput(t *testing.T) {
	mockAgent := &mockAgent{
		response: map[string]any{
			"output": "Blog post created",
			"metadata": map[string]string{
				"word_count": "500",
				"status":     "draft",
			},
		},
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()
	input := map[string]any{
		"input":        "Create a blog post about AI",
		"chat_history": []map[string]string{},
		"options": map[string]interface{}{
			"temperature": 0.7,
			"max_tokens":  1000,
		},
	}

	result, err := service.Chat(ctx, input)

	if err != nil {
		t.Fatalf("Chat失败: %v", err)
	}

	if result == nil {
		t.Fatal("结果不应该为nil")
	}
}

// mockAgent 模拟Agent用于测试
type mockAgent struct {
	response map[string]any
	err      error
}

func (m *mockAgent) Chat(ctx context.Context, input map[string]any) (map[string]any, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.response, nil
}

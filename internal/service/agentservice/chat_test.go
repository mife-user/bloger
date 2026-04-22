package agentservice

import (
	"context"
	"testing"

	"bloger/internal/domain"
)

// TestAgentService_Chat_Integration 测试聊天集成
func TestAgentService_Chat_Integration(t *testing.T) {
	// 创建一个更复杂的mock agent
	mockAgent := &advancedMockAgent{
		responses: []domain.ChatResponse{
			{Message: map[string]any{"output": "First response"}},
			{Message: map[string]any{"output": "Second response"}},
			{Message: map[string]any{"output": "Third response"}},
		},
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()

	// 测试多次对话
	for i := 0; i < 3; i++ {
		input := domain.ChatRequest{Message: map[string]any{
			"input": "Message",
		}}

		result, err := service.Chat(ctx, input)

		if err != nil {
			t.Errorf("第%d次Chat失败: %v", i+1, err)
			continue
		}

		// 验证响应内容
		expectedOutput := mockAgent.responses[i].Message["output"]
		if output, ok := result.Message["output"]; ok {
			if output != expectedOutput {
				t.Errorf("第%d次期望输出 '%v', 得到 '%v'", i+1, expectedOutput, output)
			}
		}
	}
}

// TestAgentService_Chat_ContextCancellation 测试上下文取消
func TestAgentService_Chat_ContextCancellation(t *testing.T) {
	mockAgent := &cancellableMockAgent{}

	service := NewAgentService(mockAgent)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消

	input := domain.ChatRequest{Message: map[string]any{
		"input": "Hello",
	}}

	_, err := service.Chat(ctx, input)

	// 应该返回context取消错误
	if err == nil {
		t.Error("Chat应该处理取消的context并返回错误")
	}
}

// TestAgentService_Chat_Timeout 测试超时
func TestAgentService_Chat_Timeout(t *testing.T) {
	mockAgent := &slowMockAgent{}

	service := NewAgentService(mockAgent)

	ctx, cancel := context.WithTimeout(context.Background(), 1)
	defer cancel()

	input := domain.ChatRequest{Message: map[string]any{
		"input": "Hello",
	}}

	_, err := service.Chat(ctx, input)

	// 应该返回超时错误
	if err == nil {
		t.Error("Chat应该处理超时并返回错误")
	}
}

// TestAgentService_Chat_Concurrent 测试并发调用
func TestAgentService_Chat_Concurrent(t *testing.T) {
	mockAgent := &threadSafeMockAgent{
		response: domain.ChatResponse{Message: map[string]any{"output": "concurrent response"}},
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()

	// 并发调用
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			input := domain.ChatRequest{Message: map[string]any{
				"input": "Concurrent message",
			}}

			result, err := service.Chat(ctx, input)

			if err != nil {
				t.Errorf("并发调用%d失败: %v", id, err)
			}

			// 验证响应内容
			if output, ok := result.Message["output"]; ok {
				if output != "concurrent response" {
					t.Errorf("并发调用%d期望输出 'concurrent response', 得到 '%v'", id, output)
				}
			}

			done <- true
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}
}

// advancedMockAgent 高级mock agent
type advancedMockAgent struct {
	responses []domain.ChatResponse
	callCount int
}

func (m *advancedMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	if m.callCount >= len(m.responses) {
		m.callCount = 0
	}
	response := m.responses[m.callCount]
	m.callCount++
	return response, nil
}

// cancellableMockAgent 可取消的mock agent
type cancellableMockAgent struct{}

func (m *cancellableMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	select {
	case <-ctx.Done():
		return domain.ChatResponse{}, ctx.Err()
	default:
		return domain.ChatResponse{Message: map[string]any{"output": "response"}}, nil
	}
}

// slowMockAgent 慢速mock agent
type slowMockAgent struct{}

func (m *slowMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	select {
	case <-ctx.Done():
		return domain.ChatResponse{}, ctx.Err()
	case <-make(chan bool): // 永远不会返回
		return domain.ChatResponse{}, nil
	}
}

// threadSafeMockAgent 线程安全的mock agent
type threadSafeMockAgent struct {
	response domain.ChatResponse
}

func (m *threadSafeMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	// 返回响应的副本，避免并发问题
	result := domain.ChatResponse{Message: make(map[string]any)}
	for k, v := range m.response.Message {
		result.Message[k] = v
	}
	return result, nil
}

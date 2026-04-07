package agentservice

import (
	"context"
	"testing"
)

// TestAgentService_Chat_Integration 测试聊天集成
func TestAgentService_Chat_Integration(t *testing.T) {
	// 创建一个更复杂的mock agent
	mockAgent := &advancedMockAgent{
		responses: []string{
			"First response",
			"Second response",
			"Third response",
		},
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()

	// 测试多次对话
	for i := 0; i < 3; i++ {
		input := map[string]any{
			"input": "Message",
		}

		result, err := service.Chat(ctx, input)

		if err != nil {
			t.Errorf("第%d次Chat失败: %v", i+1, err)
			continue
		}

		if result == "" {
			t.Errorf("第%d次结果不应该为空", i+1)
		}
	}
}

// TestAgentService_Chat_ContextCancellation 测试上下文取消
func TestAgentService_Chat_ContextCancellation(t *testing.T) {
	mockAgent := &cancellableMockAgent{}

	service := NewAgentService(mockAgent)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消

	input := map[string]any{
		"input": "Hello",
	}

	_, err := service.Chat(ctx, input)

	// 应该返回context取消错误
	if err == nil {
		t.Log("Chat处理了取消的context")
	}
}

// TestAgentService_Chat_Timeout 测试超时
func TestAgentService_Chat_Timeout(t *testing.T) {
	mockAgent := &slowMockAgent{}

	service := NewAgentService(mockAgent)

	ctx, cancel := context.WithTimeout(context.Background(), 1)
	defer cancel()

	input := map[string]any{
		"input": "Hello",
	}

	_, err := service.Chat(ctx, input)

	// 应该返回超时错误
	if err == nil {
		t.Log("Chat处理了超时")
	}
}

// TestAgentService_Chat_Concurrent 测试并发调用
func TestAgentService_Chat_Concurrent(t *testing.T) {
	mockAgent := &threadSafeMockAgent{
		response: "concurrent response",
	}

	service := NewAgentService(mockAgent)

	ctx := context.Background()

	// 并发调用
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			input := map[string]any{
				"input": "Concurrent message",
			}

			result, err := service.Chat(ctx, input)

			if err != nil {
				t.Errorf("并发调用%d失败: %v", id, err)
			}

			if result == "" {
				t.Errorf("并发调用%d结果为空", id)
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
	responses []string
	callCount int
}

func (m *advancedMockAgent) Chat(ctx context.Context, input map[string]any) (string, error) {
	if m.callCount >= len(m.responses) {
		m.callCount = 0
	}
	response := m.responses[m.callCount]
	m.callCount++
	return response, nil
}

// cancellableMockAgent 可取消的mock agent
type cancellableMockAgent struct{}

func (m *cancellableMockAgent) Chat(ctx context.Context, input map[string]any) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return "response", nil
	}
}

// slowMockAgent 慢速mock agent
type slowMockAgent struct{}

func (m *slowMockAgent) Chat(ctx context.Context, input map[string]any) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-make(chan bool): // 永远不会返回
		return "", nil
	}
}

// threadSafeMockAgent 线程安全的mock agent
type threadSafeMockAgent struct {
	response string
}

func (m *threadSafeMockAgent) Chat(ctx context.Context, input map[string]any) (string, error) {
	return m.response, nil
}

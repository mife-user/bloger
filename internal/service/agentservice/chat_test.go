package agentservice

import (
	"context"
	"testing"

	"mifer/internal/domain"
)

func TestAgentService_Chat_Integration(t *testing.T) {
	mockAgent := &integrationMockAgent{
		responses: []domain.ChatResponse{
			{Content: "First response"},
			{Content: "Second response"},
			{Content: "Third response"},
		},
	}
	service := NewAgentService(mockAgent)
	ctx := context.Background()
	for i, expected := range []string{"First response", "Second response", "Third response"} {
		result, err := service.Chat(ctx, domain.ChatRequest{Content: "Message"})
		if err != nil {
			t.Errorf("Chat #%d failed: %v", i+1, err)
			continue
		}
		if result.Content != expected {
			t.Errorf("Chat #%d expected '%s', got '%s'", i+1, expected, result.Content)
		}
	}
}

func TestAgentService_Chat_ContextCancellation(t *testing.T) {
	service := NewAgentService(&cancellableMockAgent{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := service.Chat(ctx, domain.ChatRequest{Content: "Hello"})
	if err == nil {
		t.Error("Chat should handle cancelled context")
	}
}

func TestAgentService_Chat_Timeout(t *testing.T) {
	service := NewAgentService(&slowMockAgent{})
	ctx, cancel := context.WithTimeout(context.Background(), 1)
	defer cancel()
	_, err := service.Chat(ctx, domain.ChatRequest{Content: "Hello"})
	if err == nil {
		t.Error("Chat should handle timeout")
	}
}

func TestAgentService_Chat_Concurrent(t *testing.T) {
	mockAgent := &concurrentMockAgent{
		response: domain.ChatResponse{Content: "concurrent response"},
	}
	service := NewAgentService(mockAgent)
	ctx := context.Background()
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			result, err := service.Chat(ctx, domain.ChatRequest{Content: "Concurrent message"})
			if err != nil {
				t.Errorf("concurrent call %d failed: %v", id, err)
			}
			if result.Content != "concurrent response" {
				t.Errorf("concurrent call %d expected 'concurrent response', got '%s'", id, result.Content)
			}
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

type integrationMockAgent struct {
	responses []domain.ChatResponse
	callCount int
}

func (m *integrationMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	if m.callCount >= len(m.responses) {
		m.callCount = 0
	}
	r := m.responses[m.callCount]
	m.callCount++
	return r, nil
}

type cancellableMockAgent struct{}

func (m *cancellableMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	select {
	case <-ctx.Done():
		return domain.ChatResponse{}, ctx.Err()
	default:
		return domain.ChatResponse{Content: "response"}, nil
	}
}

type slowMockAgent struct{}

func (m *slowMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	select {
	case <-ctx.Done():
		return domain.ChatResponse{}, ctx.Err()
	case <-make(chan bool):
		return domain.ChatResponse{}, nil
	}
}

type concurrentMockAgent struct {
	response domain.ChatResponse
}

func (m *concurrentMockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	return domain.ChatResponse{Content: m.response.Content}, nil
}

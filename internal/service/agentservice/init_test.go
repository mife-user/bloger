package agentservice

import (
	"context"
	"testing"

	"mifer/internal/domain"
)

func TestNewAgentService(t *testing.T) {
	mockAgent := &mockAgent{}
	service := NewAgentService(mockAgent)
	if service == nil {
		t.Fatal("AgentService should not be nil")
	}
	if service.agent == nil {
		t.Fatal("AgentService.agent should not be nil")
	}
}

func TestNewAgentService_NilAgent(t *testing.T) {
	service := NewAgentService(nil)
	if service == nil {
		t.Fatal("AgentService should not be nil")
	}
	if service.agent != nil {
		t.Error("agent should be nil when nil passed")
	}
}

func TestAgentService_Chat(t *testing.T) {
	mockAgent := &mockAgent{
		response: domain.ChatResponse{Content: "Hello! How can I help you?"},
	}
	service := NewAgentService(mockAgent)
	ctx := context.Background()
	result, err := service.Chat(ctx, domain.ChatRequest{Content: "Hello"})
	if err != nil {
		t.Fatalf("Chat failed: %v", err)
	}
	if result.Content != "Hello! How can I help you?" {
		t.Errorf("expected 'Hello! How can I help you?', got '%s'", result.Content)
	}
}

func TestAgentService_Chat_Error(t *testing.T) {
	mockAgent := &mockAgent{err: context.DeadlineExceeded}
	service := NewAgentService(mockAgent)
	_, err := service.Chat(context.Background(), domain.ChatRequest{Content: "Hello"})
	if err == nil {
		t.Fatal("should return error")
	}
}

func TestAgentService_Chat_EmptyInput(t *testing.T) {
	mockAgent := &mockAgent{
		response: domain.ChatResponse{Content: "Please provide input"},
	}
	service := NewAgentService(mockAgent)
	result, err := service.Chat(context.Background(), domain.ChatRequest{Content: ""})
	if err != nil {
		t.Fatalf("Chat failed: %v", err)
	}
	if result.Content != "Please provide input" {
		t.Errorf("expected 'Please provide input', got '%s'", result.Content)
	}
}

type mockAgent struct {
	response domain.ChatResponse
	err      error
}

func (m *mockAgent) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	if m.err != nil {
		return domain.ChatResponse{}, m.err
	}
	return m.response, nil
}

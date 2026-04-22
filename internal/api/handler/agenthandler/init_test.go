package agenthandler

import (
	"context"
	"testing"

	"bloger/internal/domain"
)

// TestNewAgentHandler 测试创建AgentHandler
func TestNewAgentHandler(t *testing.T) {
	// 创建mock service
	mockService := &mockAgentService{}

	handler := NewAgentHandler(mockService)

	if handler == nil {
		t.Fatal("AgentHandler不应该为nil")
	}

	if handler.service == nil {
		t.Fatal("AgentHandler.service不应该为nil")
	}
}

// TestNewAgentHandler_NilService 测试nil service
func TestNewAgentHandler_NilService(t *testing.T) {
	handler := NewAgentHandler(nil)

	// NewAgentHandler应该能处理nil service
	if handler == nil {
		t.Fatal("AgentHandler不应该为nil")
	}

	// 但service应该为nil
	if handler.service != nil {
		t.Error("传入nil时，service应该为nil")
	}
}

// TestAgentHandler_ServiceInjection 测试服务注入
func TestAgentHandler_ServiceInjection(t *testing.T) {
	mockService := &mockAgentService{}
	handler := NewAgentHandler(mockService)

	// 验证服务注入成功
	if handler.service == nil {
		t.Error("服务注入失败")
	}
}

// mockAgentService 模拟AgentService
type mockAgentService struct{}

func (m *mockAgentService) Chat(ctx context.Context, input domain.ChatRequest) (domain.ChatResponse, error) {
	return domain.ChatResponse{Message: map[string]any{"output": "test response"}}, nil
}

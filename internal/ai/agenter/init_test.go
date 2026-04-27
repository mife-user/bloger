package agenter

import (
	"context"
	"testing"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type mockChatModel struct{}

func (m *mockChatModel) Generate(ctx context.Context, input []*schema.Message, opts ...any) (*schema.Message, error) {
	return schema.AssistantMessage("mock response", nil), nil
}

func (m *mockChatModel) Stream(ctx context.Context, input []*schema.Message, opts ...any) (*schema.StreamReader[*schema.Message], error) {
	return nil, nil
}

type mockTool struct{}

func (m *mockTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{Name: "mock", Desc: "a mock tool"}, nil
}

// TestInitAgent — skipped because react.NewAgent requires a real ToolCallingChatModel
// and cannot work with our limited mock. Use integration tests with real API keys.
func TestInitAgent_Mock(t *testing.T) {
	t.Skip("requires real ToolCallingChatModel; use integration test")
}

// Verify that mock types satisfy interfaces at compile time
var _ tool.BaseTool = (*mockTool)(nil)

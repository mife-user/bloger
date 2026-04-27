package hugotool

import (
	"context"
	"testing"
)

func TestHugTool_Info(t *testing.T) {
	tool := &HugTool{}
	info, err := tool.Info(context.Background())
	if err != nil {
		t.Fatalf("Info failed: %v", err)
	}
	if info.Name != "hug" {
		t.Errorf("expected name 'hug', got '%s'", info.Name)
	}
	if info.Desc != "hug tool" {
		t.Errorf("expected desc 'hug tool', got '%s'", info.Desc)
	}
}

func TestHugTool_InvokableRun(t *testing.T) {
	tool := &HugTool{}
	result, err := tool.InvokableRun(context.Background(), "test input")
	if err != nil {
		t.Errorf("InvokableRun should not return error: %v", err)
	}
	if result != "" {
		t.Errorf("expected empty result, got '%s'", result)
	}
}

func TestHugTool_InvokableRun_EmptyInput(t *testing.T) {
	tool := &HugTool{}
	result, err := tool.InvokableRun(context.Background(), "")
	if err != nil {
		t.Errorf("empty input should not return error: %v", err)
	}
	if result != "" {
		t.Errorf("expected empty result, got '%s'", result)
	}
}

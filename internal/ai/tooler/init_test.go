package tooler

import (
	"context"
	"testing"
)

func TestInitTools(t *testing.T) {
	tools := InitTools()
	if tools == nil {
		t.Fatal("tools should not be nil")
	}
	if len(tools) != 2 {
		t.Errorf("expected 2 tools, got %d", len(tools))
	}
}

func TestInitTools_Names(t *testing.T) {
	tools := InitTools()
	names := make(map[string]bool)
	for _, tl := range tools {
		info, err := tl.Info(context.Background())
		if err != nil {
			t.Fatalf("Info() failed: %v", err)
		}
		names[info.Name] = true
	}
	if !names["git"] {
		t.Error("missing git tool")
	}
	if !names["hug"] {
		t.Error("missing hug tool")
	}
}

func TestInitTools_Descriptions(t *testing.T) {
	tools := InitTools()
	for _, tl := range tools {
		info, _ := tl.Info(context.Background())
		if info.Desc == "" {
			t.Errorf("tool %s has no description", info.Name)
		}
	}
}

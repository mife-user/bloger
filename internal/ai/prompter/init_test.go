package prompter

import (
	"context"
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestModifierBuilder_EmptyPrompt(t *testing.T) {
	modifier := ModifierBuilder{}.Build("")
	result := modifier(context.Background(), []*schema.Message{
		schema.UserMessage("hello"),
	})
	// system message prepended, then original input
	if len(result) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(result))
	}
	if result[0].Role != schema.System {
		t.Errorf("first message should be system, got %s", result[0].Role)
	}
}

func TestModifierBuilder_Prompt(t *testing.T) {
	modifier := ModifierBuilder{}.Build("You are a helpful assistant.")
	result := modifier(context.Background(), []*schema.Message{
		schema.UserMessage("hi"),
	})
	if result[0].Content != "You are a helpful assistant." {
		t.Errorf("expected system prompt, got '%s'", result[0].Content)
	}
}

func TestModifierBuilder_MultilinePrompt(t *testing.T) {
	prompt := "You are an assistant\nwith multiple\nlines."
	modifier := ModifierBuilder{}.Build(prompt)
	result := modifier(context.Background(), nil)
	if len(result) != 1 {
		t.Errorf("expected 1 message, got %d", len(result))
	}
}

func TestModifierBuilder_SpecialCharacters(t *testing.T) {
	modifier := ModifierBuilder{}.Build("你是一个助手 with emoji 🤖")
	result := modifier(context.Background(), []*schema.Message{
		schema.UserMessage("中文输入"),
	})
	if len(result) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(result))
	}
}

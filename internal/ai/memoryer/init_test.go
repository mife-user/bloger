package memoryer

import (
	"testing"

	"github.com/cloudwego/eino/schema"
)

func TestChatHistory_Append(t *testing.T) {
	h := New(2048)
	h.Append(schema.UserMessage("hello"))
	h.Append(schema.AssistantMessage("hi", nil))
	if len(h.Messages()) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(h.Messages()))
	}
}

func TestChatHistory_Truncation(t *testing.T) {
	h := New(1) // 1 token = 4 chars
	h.Append(schema.UserMessage("12345")) // 5 chars > 4, but 1 msg won't trim
	h.Append(schema.AssistantMessage("ok", nil)) // now 2 msgs, oldest trimmed
	if len(h.Messages()) != 1 {
		t.Fatalf("expected 1 message after truncation, got %d", len(h.Messages()))
	}
}

func TestChatHistory_Messages(t *testing.T) {
	h := New(2048)
	if len(h.Messages()) != 0 {
		t.Error("expected empty history")
	}
	h.Append(schema.UserMessage("test"))
	if len(h.Messages()) != 1 {
		t.Error("expected 1 message")
	}
}

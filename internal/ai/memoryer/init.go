package memoryer

import "github.com/cloudwego/eino/schema"

type ChatHistory struct {
	messages  []*schema.Message
	maxTokens int
}

func New(maxTokens int) *ChatHistory {
	return &ChatHistory{
		messages:  make([]*schema.Message, 0, maxTokens),
		maxTokens: maxTokens,
	}
}

func (h *ChatHistory) Append(msg *schema.Message) {
	h.messages = append(h.messages, msg)
	for h.charCount() > h.maxTokens*4 && len(h.messages) > 1 {
		h.messages = append(h.messages[:1], h.messages[2:]...)
	}
}

func (h *ChatHistory) Messages() []*schema.Message {
	return h.messages
}

func (h *ChatHistory) charCount() int {
	n := 0
	for _, m := range h.messages {
		n += len(m.Content)
	}
	return n
}

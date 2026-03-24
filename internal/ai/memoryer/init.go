package memoryer

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
)

// InitMemoryer 初始化内存管理器
func InitMemoryer(llm llms.Model) *memory.ConversationTokenBuffer {
	return memory.NewConversationTokenBuffer(llm, 10)
}

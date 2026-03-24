package tooler

import (
	"bloger/internal/ai/tooler/gittool"
	"bloger/internal/ai/tooler/hugotool"

	"github.com/tmc/langchaingo/tools"
)

// InitTools 初始化工具
func InitTools() []tools.Tool {
	return []tools.Tool{
		&hugotool.HugTool{},
		&gittool.GitTool{},
	}
}

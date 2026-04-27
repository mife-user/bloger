package tooler

import (
	"mifer/internal/ai/tooler/gittool"
	"mifer/internal/ai/tooler/hugotool"

	"github.com/cloudwego/eino/components/tool"
)

func InitTools() []tool.BaseTool {
	return []tool.BaseTool{
		&hugotool.HugTool{},
		&gittool.GitTool{},
	}
}

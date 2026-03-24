package hugotool

import "context"

type HugTool struct {
}

// Name 工具名称
func (h *HugTool) Name() string {
	return "hug"
}

// Description 工具描述
func (h *HugTool) Description() string {
	return "hug tool"
}

// Call 调用工具
func (h *HugTool) Call(ctx context.Context, input string) (string, error) {
	return "", nil
}

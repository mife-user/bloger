package gittool

import "context"

type GitTool struct {
}

// Name 工具名称
func (g *GitTool) Name() string {
	return "git"
}

// Description 工具描述
func (g *GitTool) Description() string {
	return "git tool"
}

// Call 调用工具
func (g *GitTool) Call(ctx context.Context, input string) (string, error) {
	return "", nil
}

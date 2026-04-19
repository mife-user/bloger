package gittool

import (
	"bloger/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
)

// status 查看仓库状态
func (g *GitTool) status(input GitInput) *GitOutput {
	path := input.Path
	if path == "" {
		return &GitOutput{
			Success: false,
			Message: "仓库路径不能为空",
		}
	}

	// 打开仓库
	repo, err := git.PlainOpen(path)
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("打开仓库失败: %v", err),
		}
	}

	// 获取工作树
	worktree, err := repo.Worktree()
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("获取工作树失败: %v", err),
		}
	}

	// 获取状态
	status, err := worktree.Status()
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("获取状态失败: %v", err),
		}
	}

	// 格式化状态信息
	var changes []map[string]string
	for file, s := range status {
		changes = append(changes, map[string]string{
			"file":     file,
			"staging":  string(rune(s.Staging)),
			"worktree": string(rune(s.Worktree)),
		})
	}

	logger.Info(fmt.Sprintf("仓库状态: %d 个文件有变更", len(changes)))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("仓库状态: %d 个文件有变更", len(changes)),
		Data: map[string]any{
			"changes": changes,
		},
	}
}

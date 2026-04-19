package gittool

import (
	"bloger/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
)

// commit 提交更改
func (g *GitTool) commit(input GitInput) *GitOutput {
	path := input.Path
	if path == "" {
		return &GitOutput{
			Success: false,
			Message: "仓库路径不能为空",
		}
	}

	if input.Message == "" {
		return &GitOutput{
			Success: false,
			Message: "提交信息不能为空",
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

	// 提交
	commit, err := worktree.Commit(input.Message, &git.CommitOptions{
		Author: g.author,
	})
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("提交失败: %v", err),
		}
	}

	logger.Info(fmt.Sprintf("成功提交: %s (hash: %s)", input.Message, commit.String()))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("成功提交: %s", input.Message),
		Data: map[string]string{
			"commit_hash": commit.String(),
		},
	}
}

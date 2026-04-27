package gittool

import (
	"mifer/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
)

// addFiles 添加文件到暂存区
func (g *GitTool) addFiles(input GitInput) *GitOutput {
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

	// 添加文件
	if len(input.Files) == 0 {
		// 如果没有指定文件，添加所有文件
		if err := worktree.AddWithOptions(&git.AddOptions{
			All: true,
		}); err != nil {
			return &GitOutput{
				Success: false,
				Message: fmt.Sprintf("添加文件失败: %v", err),
			}
		}
		logger.Info("已添加所有文件到暂存区")
	} else {
		// 添加指定文件
		for _, file := range input.Files {
			if _, err := worktree.Add(file); err != nil {
				return &GitOutput{
					Success: false,
					Message: fmt.Sprintf("添加文件 %s 失败: %v", file, err),
				}
			}
			logger.Info(fmt.Sprintf("已添加文件: %s", file))
		}
	}

	return &GitOutput{
		Success: true,
		Message: "成功添加文件到暂存区",
		Data: map[string]any{
			"files": input.Files,
		},
	}
}

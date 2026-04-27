package gittool

import (
	"mifer/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// checkout 切换分支
func (g *GitTool) checkout(input GitInput) *GitOutput {
	path := input.Path
	if path == "" {
		return &GitOutput{
			Success: false,
			Message: "仓库路径不能为空",
		}
	}

	if input.BranchName == "" {
		return &GitOutput{
			Success: false,
			Message: "分支名称不能为空",
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

	// 切换分支
	if err := worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/heads/" + input.BranchName),
		Create: input.CreateBranch,
	}); err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("切换分支失败: %v", err),
		}
	}

	logger.Info(fmt.Sprintf("成功切换到分支: %s", input.BranchName))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("成功切换到分支: %s", input.BranchName),
	}
}

package gittool

import (
	"bloger/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// branch 分支管理
func (g *GitTool) branch(input GitInput) *GitOutput {
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

	if input.CreateBranch && input.BranchName != "" {
		// 创建新分支
		head, err := repo.Head()
		if err != nil {
			return &GitOutput{
				Success: false,
				Message: fmt.Sprintf("获取HEAD失败: %v", err),
			}
		}

		ref := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+input.BranchName), head.Hash())
		if err := repo.Storer.SetReference(ref); err != nil {
			return &GitOutput{
				Success: false,
				Message: fmt.Sprintf("创建分支失败: %v", err),
			}
		}

		logger.Info(fmt.Sprintf("成功创建分支: %s", input.BranchName))

		return &GitOutput{
			Success: true,
			Message: fmt.Sprintf("成功创建分支: %s", input.BranchName),
		}
	}

	// 列出所有分支
	branches, err := repo.Branches()
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("获取分支列表失败: %v", err),
		}
	}

	var branchList []string
	if err := branches.ForEach(func(ref *plumbing.Reference) error {
		branchList = append(branchList, ref.Name().Short())
		return nil
	}); err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("遍历分支失败: %v", err),
		}
	}

	logger.Info(fmt.Sprintf("当前分支列表: %v", branchList))

	return &GitOutput{
		Success: true,
		Message: "获取分支列表成功",
		Data: map[string]any{
			"branches": branchList,
		},
	}
}

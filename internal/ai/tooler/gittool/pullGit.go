package gittool

import (
	"bloger/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// pull 拉取远程更新
func (g *GitTool) pull(input GitInput) *GitOutput {
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

	// 设置认证信息
	var auth *http.BasicAuth
	if input.Token != "" {
		auth = &http.BasicAuth{
			Username: "oauth2",
			Password: input.Token,
		}
	}

	// 拉取
	remoteName := input.RemoteName
	if remoteName == "" {
		remoteName = "origin"
	}

	if err := worktree.Pull(&git.PullOptions{
		RemoteName: remoteName,
		Auth:       auth,
	}); err != nil && err != git.NoErrAlreadyUpToDate {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("拉取失败: %v", err),
		}
	}

	logger.Info(fmt.Sprintf("成功拉取远程更新: %s", remoteName))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("成功拉取远程更新: %s", remoteName),
	}
}

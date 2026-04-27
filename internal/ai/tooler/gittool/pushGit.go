package gittool

import (
	"mifer/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// push 推送到远程仓库
func (g *GitTool) push(input GitInput) *GitOutput {
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

	// 设置认证信息
	var auth *http.BasicAuth
	if input.Token != "" {
		auth = &http.BasicAuth{
			Username: "oauth2",
			Password: input.Token,
		}
	}

	// 推送
	remoteName := input.RemoteName
	if remoteName == "" {
		remoteName = "origin"
	}

	if err := repo.Push(&git.PushOptions{
		RemoteName: remoteName,
		Auth:       auth,
	}); err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("推送失败: %v", err),
		}
	}

	logger.Info(fmt.Sprintf("成功推送到远程仓库: %s", remoteName))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("成功推送到远程仓库: %s", remoteName),
	}
}

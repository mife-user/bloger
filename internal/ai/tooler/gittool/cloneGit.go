package gittool

import (
	"mifer/pkg/logger"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// clone 克隆远程仓库
func (g *GitTool) clone(input GitInput) *GitOutput {
	if input.RemoteURL == "" {
		return &GitOutput{
			Success: false,
			Message: "远程仓库URL不能为空",
		}
	}

	path := input.Path
	if path == "" {
		return &GitOutput{
			Success: false,
			Message: "目标路径不能为空",
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

	// 克隆仓库
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:  input.RemoteURL,
		Auth: auth,
	})
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("克隆失败: %v", err),
		}
	}

	logger.Info(fmt.Sprintf("成功克隆仓库: %s -> %s", input.RemoteURL, path))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("成功克隆仓库: %s -> %s", input.RemoteURL, path),
		Data: map[string]string{
			"remote_url": input.RemoteURL,
			"path":       path,
		},
	}
}

package gittool

import (
	"mifer/pkg/logger"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

// remote 远程仓库管理
func (g *GitTool) remote(input GitInput) *GitOutput {
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

	if input.RemoteURL != "" && input.RemoteName != "" {
		// 添加远程仓库
		_, err := repo.CreateRemote(&config.RemoteConfig{
			Name: input.RemoteName,
			URLs: []string{input.RemoteURL},
		})
		if err != nil {
			return &GitOutput{
				Success: false,
				Message: fmt.Sprintf("添加远程仓库失败: %v", err),
			}
		}

		logger.Info(fmt.Sprintf("成功添加远程仓库: %s -> %s", input.RemoteName, input.RemoteURL))

		return &GitOutput{
			Success: true,
			Message: fmt.Sprintf("成功添加远程仓库: %s -> %s", input.RemoteName, input.RemoteURL),
		}
	}

	// 列出所有远程仓库
	remotes, err := repo.Remotes()
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("获取远程仓库列表失败: %v", err),
		}
	}

	var remoteList []map[string]string
	for _, remote := range remotes {
		remoteList = append(remoteList, map[string]string{
			"name": remote.Config().Name,
			"urls": strings.Join(remote.Config().URLs, ", "),
		})
	}

	logger.Info(fmt.Sprintf("远程仓库列表: %v", remoteList))

	return &GitOutput{
		Success: true,
		Message: "获取远程仓库列表成功",
		Data: map[string]any{
			"remotes": remoteList,
		},
	}
}

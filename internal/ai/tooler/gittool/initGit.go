package gittool

import (
	"bloger/pkg/logger"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// initRepository 初始化Git仓库
func (g *GitTool) initRepository(input GitInput) *GitOutput {
	path := input.Path
	if path == "" {
		return &GitOutput{
			Success: false,
			Message: "仓库路径不能为空",
		}
	}

	// 确保目录存在
	if err := os.MkdirAll(path, 0755); err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("创建目录失败: %v", err),
		}
	}

	// 初始化仓库
	_, err := git.PlainInit(path, false)
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("初始化仓库失败: %v", err),
		}
	}

	logger.Info(fmt.Sprintf("成功初始化Git仓库: %s", path))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("成功初始化Git仓库: %s", path),
		Data: map[string]string{
			"path": path,
		},
	}
}

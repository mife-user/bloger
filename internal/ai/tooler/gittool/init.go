package gittool

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing/object"
)

// GitOperation 定义Git操作类型
type GitOperation string

const (
	OpInit     GitOperation = "init"     // 初始化仓库
	OpAdd      GitOperation = "add"      // 添加文件
	OpCommit   GitOperation = "commit"   // 提交更改
	OpPush     GitOperation = "push"     // 推送到远程
	OpPull     GitOperation = "pull"     // 拉取更新
	OpClone    GitOperation = "clone"    // 克隆仓库
	OpBranch   GitOperation = "branch"   // 分支管理
	OpStatus   GitOperation = "status"   // 状态查询
	OpRemote   GitOperation = "remote"   // 远程仓库管理
	OpCheckout GitOperation = "checkout" // 切换分支
)

// GitInput 定义Git工具的输入参数
type GitInput struct {
	Operation    GitOperation `json:"operation"`     // 操作类型
	Path         string       `json:"path"`          // 仓库路径
	Message      string       `json:"message"`       // 提交信息
	RemoteName   string       `json:"remote_name"`   // 远程仓库名称
	RemoteURL    string       `json:"remote_url"`    // 远程仓库URL
	BranchName   string       `json:"branch_name"`   // 分支名称
	Token        string       `json:"token"`         // GitHub token
	Username     string       `json:"username"`      // 用户名
	Email        string       `json:"email"`         // 邮箱
	Files        []string     `json:"files"`         // 文件列表
	CreateBranch bool         `json:"create_branch"` // 是否创建新分支
}

// GitOutput 定义Git工具的输出
type GitOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// GitTool Git工具结构体
type GitTool struct {
	author *object.Signature
}

// NewGitTool 创建Git工具实例
func NewGitTool(username, email string) *GitTool {
	return &GitTool{
		author: &object.Signature{
			Name:  username,
			Email: email,
		},
	}
}

// Name 工具名称
func (g *GitTool) Name() string {
	return "git"
}

// Description 工具描述
func (g *GitTool) Description() string {
	return `Git版本控制工具，支持以下操作：
- init: 初始化Git仓库
- add: 添加文件到暂存区
- commit: 提交更改
- push: 推送到远程仓库
- pull: 拉取远程更新
- clone: 克隆远程仓库
- branch: 分支管理
- status: 查看仓库状态
- remote: 远程仓库管理
- checkout: 切换分支

输入JSON格式：
{
  "operation": "操作类型",
  "path": "仓库路径",
  "message": "提交信息",
  "remote_name": "远程仓库名称",
  "remote_url": "远程仓库URL",
  "branch_name": "分支名称",
  "token": "GitHub token",
  "username": "用户名",
  "email": "邮箱",
  "files": ["文件列表"],
  "create_branch": true/false
}`
}

// Call 调用工具
func (g *GitTool) Call(ctx context.Context, input string) (string, error) {
	// 解析输入参数
	var gitInput GitInput
	if err := json.Unmarshal([]byte(input), &gitInput); err != nil {
		return g.errorOutput(fmt.Sprintf("解析输入参数失败: %v", err)), nil
	}

	// 设置默认作者信息
	if g.author == nil {
		g.author = &object.Signature{
			Name:  gitInput.Username,
			Email: gitInput.Email,
		}
	}

	// 根据操作类型执行相应操作
	var result *GitOutput
	switch gitInput.Operation {
	case OpInit:
		result = g.initRepository(gitInput)
	case OpAdd:
		result = g.addFiles(gitInput)
	case OpCommit:
		result = g.commit(gitInput)
	case OpPush:
		result = g.push(gitInput)
	case OpPull:
		result = g.pull(gitInput)
	case OpClone:
		result = g.clone(gitInput)
	case OpBranch:
		result = g.branch(gitInput)
	case OpStatus:
		result = g.status(gitInput)
	case OpRemote:
		result = g.remote(gitInput)
	case OpCheckout:
		result = g.checkout(gitInput)
	default:
		result = &GitOutput{
			Success: false,
			Message: fmt.Sprintf("不支持的操作类型: %s", gitInput.Operation),
		}
	}

	// 返回JSON格式结果
	output, err := json.Marshal(result)
	if err != nil {
		return g.errorOutput(fmt.Sprintf("序列化结果失败: %v", err)), nil
	}

	return string(output), nil
}

// errorOutput 生成错误输出
func (g *GitTool) errorOutput(message string) string {
	output := GitOutput{
		Success: false,
		Message: message,
	}
	result, _ := json.Marshal(output)
	return string(result)
}

// QuickInit 快速初始化仓库并创建初始提交
func (g *GitTool) QuickInit(path, message string) error {
	// 初始化仓库
	result := g.initRepository(GitInput{Path: path})
	if !result.Success {
		return fmt.Errorf("初始化仓库失败: %s", result.Message)
	}

	// 创建README文件
	readmePath := filepath.Join(path, "README.md")
	if err := os.WriteFile(readmePath, []byte("# My Blog\n"), 0644); err != nil {
		return fmt.Errorf("创建README失败: %v", err)
	}

	// 添加文件
	result = g.addFiles(GitInput{Path: path})
	if !result.Success {
		return fmt.Errorf("添加文件失败: %s", result.Message)
	}

	// 提交
	if message == "" {
		message = "Initial commit"
	}
	result = g.commit(GitInput{Path: path, Message: message})
	if !result.Success {
		return fmt.Errorf("提交失败: %s", result.Message)
	}

	return nil
}

// QuickPush 快速推送（添加、提交、推送一条龙）
func (g *GitTool) QuickPush(path, message, token string) error {
	// 添加所有文件
	result := g.addFiles(GitInput{Path: path})
	if !result.Success {
		return fmt.Errorf("添加文件失败: %s", result.Message)
	}

	// 提交
	result = g.commit(GitInput{Path: path, Message: message})
	if !result.Success {
		return fmt.Errorf("提交失败: %s", result.Message)
	}

	// 推送
	result = g.push(GitInput{Path: path, Token: token})
	if !result.Success {
		return fmt.Errorf("推送失败: %s", result.Message)
	}

	return nil
}

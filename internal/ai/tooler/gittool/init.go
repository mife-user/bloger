package gittool

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"bloger/pkg/logger"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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

// status 查看仓库状态
func (g *GitTool) status(input GitInput) *GitOutput {
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

	// 获取状态
	status, err := worktree.Status()
	if err != nil {
		return &GitOutput{
			Success: false,
			Message: fmt.Sprintf("获取状态失败: %v", err),
		}
	}

	// 格式化状态信息
	var changes []map[string]string
	for file, s := range status {
		changes = append(changes, map[string]string{
			"file":     file,
			"staging":  string(rune(s.Staging)),
			"worktree": string(rune(s.Worktree)),
		})
	}

	logger.Info(fmt.Sprintf("仓库状态: %d 个文件有变更", len(changes)))

	return &GitOutput{
		Success: true,
		Message: fmt.Sprintf("仓库状态: %d 个文件有变更", len(changes)),
		Data: map[string]any{
			"changes": changes,
		},
	}
}

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

package gittool

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestGitTool_Name 测试工具名称
func TestGitTool_Name(t *testing.T) {
	tool := &GitTool{}

	name := tool.Name()

	if name != "git" {
		t.Errorf("期望名称 'git', 得到 '%s'", name)
	}
}

// TestGitTool_Description 测试工具描述
func TestGitTool_Description(t *testing.T) {
	tool := &GitTool{}

	desc := tool.Description()

	if desc == "" {
		t.Error("描述不应该为空")
	}

	if !contains(desc, "Git版本控制工具") {
		t.Error("描述应该包含 'Git版本控制工具'")
	}
}

// TestGitTool_Init 测试初始化仓库
func TestGitTool_Init(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	// 创建临时目录
	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	input := GitInput{
		Operation: OpInit,
		Path:      repoPath,
	}

	inputJSON, _ := json.Marshal(input)
	result, err := tool.Call(context.Background(), string(inputJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if !output.Success {
		t.Errorf("初始化仓库失败: %s", output.Message)
	}

	// 验证.git目录存在
	gitDir := filepath.Join(repoPath, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		t.Error(".git目录应该存在")
	}
}

// TestGitTool_Add 测试添加文件
func TestGitTool_Add(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	// 创建临时仓库
	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	// 初始化仓库
	initInput := GitInput{Operation: OpInit, Path: repoPath}
	initJSON, _ := json.Marshal(initInput)
	tool.Call(context.Background(), string(initJSON))

	// 创建测试文件
	testFile := filepath.Join(repoPath, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	// 添加文件
	addInput := GitInput{
		Operation: OpAdd,
		Path:      repoPath,
		Files:     []string{"test.txt"},
	}
	addJSON, _ := json.Marshal(addInput)
	result, err := tool.Call(context.Background(), string(addJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if !output.Success {
		t.Errorf("添加文件失败: %s", output.Message)
	}
}

// TestGitTool_Commit 测试提交
func TestGitTool_Commit(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	// 创建临时仓库
	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	// 初始化仓库
	initInput := GitInput{Operation: OpInit, Path: repoPath}
	initJSON, _ := json.Marshal(initInput)
	tool.Call(context.Background(), string(initJSON))

	// 创建并添加文件
	testFile := filepath.Join(repoPath, "test.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	addInput := GitInput{Operation: OpAdd, Path: repoPath}
	addJSON, _ := json.Marshal(addInput)
	tool.Call(context.Background(), string(addJSON))

	// 提交
	commitInput := GitInput{
		Operation: OpCommit,
		Path:      repoPath,
		Message:   "Test commit",
	}
	commitJSON, _ := json.Marshal(commitInput)
	result, err := tool.Call(context.Background(), string(commitJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if !output.Success {
		t.Errorf("提交失败: %s", output.Message)
	}
}

// TestGitTool_Status 测试状态查询
func TestGitTool_Status(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	// 创建临时仓库
	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	// 初始化仓库
	initInput := GitInput{Operation: OpInit, Path: repoPath}
	initJSON, _ := json.Marshal(initInput)
	tool.Call(context.Background(), string(initJSON))

	// 查询状态
	statusInput := GitInput{
		Operation: OpStatus,
		Path:      repoPath,
	}
	statusJSON, _ := json.Marshal(statusInput)
	result, err := tool.Call(context.Background(), string(statusJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if !output.Success {
		t.Errorf("查询状态失败: %s", output.Message)
	}
}

// TestGitTool_Branch 测试分支管理
func TestGitTool_Branch(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	// 创建临时仓库
	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	// 初始化仓库并创建初始提交
	initInput := GitInput{Operation: OpInit, Path: repoPath}
	initJSON, _ := json.Marshal(initInput)
	tool.Call(context.Background(), string(initJSON))

	// 创建并提交文件
	testFile := filepath.Join(repoPath, "test.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	addInput := GitInput{Operation: OpAdd, Path: repoPath}
	addJSON, _ := json.Marshal(addInput)
	tool.Call(context.Background(), string(addJSON))

	commitInput := GitInput{Operation: OpCommit, Path: repoPath, Message: "Initial commit"}
	commitJSON, _ := json.Marshal(commitInput)
	tool.Call(context.Background(), string(commitJSON))

	// 创建新分支
	branchInput := GitInput{
		Operation:    OpBranch,
		Path:         repoPath,
		BranchName:   "test-branch",
		CreateBranch: true,
	}
	branchJSON, _ := json.Marshal(branchInput)
	result, err := tool.Call(context.Background(), string(branchJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if !output.Success {
		t.Errorf("创建分支失败: %s", output.Message)
	}
}

// TestGitTool_InvalidOperation 测试无效操作
func TestGitTool_InvalidOperation(t *testing.T) {
	tool := &GitTool{}

	input := GitInput{
		Operation: "invalid",
	}

	inputJSON, _ := json.Marshal(input)
	result, err := tool.Call(context.Background(), string(inputJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if output.Success {
		t.Error("无效操作应该返回失败")
	}
}

// TestGitTool_InvalidJSON 测试无效JSON
func TestGitTool_InvalidJSON(t *testing.T) {
	tool := &GitTool{}

	result, err := tool.Call(context.Background(), "invalid json")

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if output.Success {
		t.Error("无效JSON应该返回失败")
	}
}

// TestGitTool_QuickInit 测试快速初始化
func TestGitTool_QuickInit(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "quick-repo")

	err := tool.QuickInit(repoPath, "Initial commit")

	if err != nil {
		t.Errorf("QuickInit失败: %v", err)
	}

	// 验证.git目录存在
	gitDir := filepath.Join(repoPath, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		t.Error(".git目录应该存在")
	}

	// 验证README.md存在
	readmePath := filepath.Join(repoPath, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		t.Error("README.md应该存在")
	}
}

// TestGitTool_EmptyPath 测试空路径
func TestGitTool_EmptyPath(t *testing.T) {
	tool := &GitTool{}

	input := GitInput{
		Operation: OpInit,
		Path:      "",
	}

	inputJSON, _ := json.Marshal(input)
	result, err := tool.Call(context.Background(), string(inputJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if output.Success {
		t.Error("空路径应该返回失败")
	}
}

// TestGitTool_EmptyMessage 测试空提交信息
func TestGitTool_EmptyMessage(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	// 初始化仓库
	initInput := GitInput{Operation: OpInit, Path: repoPath}
	initJSON, _ := json.Marshal(initInput)
	tool.Call(context.Background(), string(initJSON))

	// 尝试提交空信息
	commitInput := GitInput{
		Operation: OpCommit,
		Path:      repoPath,
		Message:   "",
	}
	commitJSON, _ := json.Marshal(commitInput)
	result, err := tool.Call(context.Background(), string(commitJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if output.Success {
		t.Error("空提交信息应该返回失败")
	}
}

// TestGitTool_Remote 测试远程仓库管理
func TestGitTool_Remote(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	// 初始化仓库
	initInput := GitInput{Operation: OpInit, Path: repoPath}
	initJSON, _ := json.Marshal(initInput)
	tool.Call(context.Background(), string(initJSON))

	// 添加远程仓库
	remoteInput := GitInput{
		Operation:  OpRemote,
		Path:       repoPath,
		RemoteName: "origin",
		RemoteURL:  "https://github.com/test/test.git",
	}
	remoteJSON, _ := json.Marshal(remoteInput)
	result, err := tool.Call(context.Background(), string(remoteJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if !output.Success {
		t.Errorf("添加远程仓库失败: %s", output.Message)
	}
}

// TestGitTool_Checkout 测试切换分支
func TestGitTool_Checkout(t *testing.T) {
	tool := NewGitTool("testuser", "test@example.com")

	tempDir := t.TempDir()
	repoPath := filepath.Join(tempDir, "test-repo")

	// 初始化仓库并创建初始提交
	initInput := GitInput{Operation: OpInit, Path: repoPath}
	initJSON, _ := json.Marshal(initInput)
	tool.Call(context.Background(), string(initJSON))

	// 创建并提交文件
	testFile := filepath.Join(repoPath, "test.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	addInput := GitInput{Operation: OpAdd, Path: repoPath}
	addJSON, _ := json.Marshal(addInput)
	tool.Call(context.Background(), string(addJSON))

	commitInput := GitInput{Operation: OpCommit, Path: repoPath, Message: "Initial commit"}
	commitJSON, _ := json.Marshal(commitInput)
	tool.Call(context.Background(), string(commitJSON))

	// 创建新分支
	branchInput := GitInput{
		Operation:    OpBranch,
		Path:         repoPath,
		BranchName:   "test-branch",
		CreateBranch: true,
	}
	branchJSON, _ := json.Marshal(branchInput)
	tool.Call(context.Background(), string(branchJSON))

	// 切换分支
	checkoutInput := GitInput{
		Operation:  OpCheckout,
		Path:       repoPath,
		BranchName: "test-branch",
	}
	checkoutJSON, _ := json.Marshal(checkoutInput)
	result, err := tool.Call(context.Background(), string(checkoutJSON))

	if err != nil {
		t.Fatalf("Call失败: %v", err)
	}

	var output GitOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("解析结果失败: %v", err)
	}

	if !output.Success {
		t.Errorf("切换分支失败: %s", output.Message)
	}
}

// 辅助函数
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

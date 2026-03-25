package gittool

import (
	"context"
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

	if desc != "git tool" {
		t.Errorf("期望描述 'git tool', 得到 '%s'", desc)
	}
}

// TestGitTool_Call 测试工具调用
func TestGitTool_Call(t *testing.T) {
	tool := &GitTool{}
	ctx := context.Background()

	result, err := tool.Call(ctx, "test input")

	// 当前实现返回空字符串和nil错误
	if err != nil {
		t.Errorf("调用不应该返回错误: %v", err)
	}

	// 当前实现返回空字符串
	if result != "" {
		t.Errorf("期望空结果, 得到 '%s'", result)
	}
}

// TestGitTool_Call_EmptyInput 测试空输入
func TestGitTool_Call_EmptyInput(t *testing.T) {
	tool := &GitTool{}
	ctx := context.Background()

	result, err := tool.Call(ctx, "")

	// 应该能处理空输入
	if err != nil {
		t.Errorf("空输入不应该返回错误: %v", err)
	}

	if result != "" {
		t.Errorf("期望空结果, 得到 '%s'", result)
	}
}

// TestGitTool_Call_NilContext 测试nil上下文
func TestGitTool_Call_NilContext(t *testing.T) {
	tool := &GitTool{}

	// 注意：实际使用中应该避免nil context
	// 但测试应该验证工具是否能处理
	result, err := tool.Call(nil, "test")

	// 当前实现不使用context，所以不会panic
	if err != nil {
		t.Errorf("不应该返回错误: %v", err)
	}

	if result != "" {
		t.Errorf("期望空结果, 得到 '%s'", result)
	}
}

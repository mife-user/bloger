package hugotool

import (
	"context"
	"testing"
)

// TestHugTool_Name 测试工具名称
func TestHugTool_Name(t *testing.T) {
	tool := &HugTool{}

	name := tool.Name()

	if name != "hug" {
		t.Errorf("期望名称 'hug', 得到 '%s'", name)
	}
}

// TestHugTool_Description 测试工具描述
func TestHugTool_Description(t *testing.T) {
	tool := &HugTool{}

	desc := tool.Description()

	if desc == "" {
		t.Error("描述不应该为空")
	}

	if desc != "hug tool" {
		t.Errorf("期望描述 'hug tool', 得到 '%s'", desc)
	}
}

// TestHugTool_Call 测试工具调用
func TestHugTool_Call(t *testing.T) {
	tool := &HugTool{}
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

// TestHugTool_Call_EmptyInput 测试空输入
func TestHugTool_Call_EmptyInput(t *testing.T) {
	tool := &HugTool{}
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

// TestHugTool_Call_DifferentInputs 测试不同输入
func TestHugTool_Call_DifferentInputs(t *testing.T) {
	tool := &HugTool{}
	ctx := context.Background()

	inputs := []string{
		"create blog post",
		"generate markdown",
		"build site",
		"中文输入",
		"emoji 🎉",
	}

	for _, input := range inputs {
		result, err := tool.Call(ctx, input)

		if err != nil {
			t.Errorf("输入 '%s' 不应该返回错误: %v", input, err)
		}

		// 当前实现返回空字符串
		if result != "" {
			t.Errorf("输入 '%s' 期望空结果, 得到 '%s'", input, result)
		}
	}
}

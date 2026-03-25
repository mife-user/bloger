package tooler

import (
	"testing"
)

// TestInitTools 测试初始化工具列表
func TestInitTools(t *testing.T) {
	tools := InitTools()

	if tools == nil {
		t.Fatal("工具列表不应该为nil")
	}

	// 应该有2个工具
	if len(tools) != 2 {
		t.Errorf("期望2个工具, 得到 %d", len(tools))
	}

	// 验证工具名称
	toolNames := make(map[string]bool)
	for _, tool := range tools {
		toolNames[tool.Name()] = true
	}

	if !toolNames["git"] {
		t.Error("缺少git工具")
	}

	if !toolNames["hug"] {
		t.Error("缺少hug工具")
	}
}

// TestInitTools_ToolDescriptions 测试工具描述
func TestInitTools_ToolDescriptions(t *testing.T) {
	tools := InitTools()

	for _, tool := range tools {
		// 每个工具应该有描述
		if tool.Description() == "" {
			t.Errorf("工具 %s 缺少描述", tool.Name())
		}
	}
}

// TestGitTool_Name 测试Git工具名称
func TestGitTool_Name(t *testing.T) {
	tools := InitTools()
	var gitTool interface {
		Name() string
	}

	for _, tool := range tools {
		if tool.Name() == "git" {
			gitTool = tool
			break
		}
	}

	if gitTool == nil {
		t.Fatal("找不到git工具")
	}

	if gitTool.Name() != "git" {
		t.Errorf("期望工具名称 'git', 得到 '%s'", gitTool.Name())
	}
}

// TestHugTool_Name 测试Hug工具名称
func TestHugTool_Name(t *testing.T) {
	tools := InitTools()
	var hugTool interface {
		Name() string
	}

	for _, tool := range tools {
		if tool.Name() == "hug" {
			hugTool = tool
			break
		}
	}

	if hugTool == nil {
		t.Fatal("找不到hug工具")
	}

	if hugTool.Name() != "hug" {
		t.Errorf("期望工具名称 'hug', 得到 '%s'", hugTool.Name())
	}
}

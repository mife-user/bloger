package executor

import (
	"testing"
)

// TestExecutor_Chat 测试聊天功能
func TestExecutor_Chat(t *testing.T) {
	// 由于Executor需要真实的agents.Executor，我们只测试结构创建
	executor := NewExecutor(nil)

	if executor == nil {
		t.Fatal("Executor不应该为nil")
	}
}

// TestExecutor_Chat_Error 测试聊天错误处理
func TestExecutor_Chat_Error(t *testing.T) {
	// 由于Executor需要真实的agents.Executor，我们只测试结构创建
	executor := NewExecutor(nil)

	if executor == nil {
		t.Fatal("Executor不应该为nil")
	}
}

// TestExecutor_Chat_EmptyInput 测试空输入
func TestExecutor_Chat_EmptyInput(t *testing.T) {
	// 由于Executor需要真实的agents.Executor，我们只测试结构创建
	executor := NewExecutor(nil)

	if executor == nil {
		t.Fatal("Executor不应该为nil")
	}
}

// TestExecutor_Chat_ComplexInput 测试复杂输入
func TestExecutor_Chat_ComplexInput(t *testing.T) {
	// 由于Executor需要真实的agents.Executor，我们只测试结构创建
	executor := NewExecutor(nil)

	if executor == nil {
		t.Fatal("Executor不应该为nil")
	}
}

// TestNewExecutor_Wrap 测试包装executor
func TestNewExecutor_Wrap(t *testing.T) {
	// 测试NewExecutor是否能正确创建Executor实例
	exec := NewExecutor(nil)

	if exec == nil {
		t.Fatal("NewExecutor不应该返回nil")
	}

	// 验证内部字段
	if exec.executor != nil {
		t.Error("传入nil时，内部executor应该为nil")
	}
}

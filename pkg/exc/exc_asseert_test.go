package exc

import (
	"testing"
)

// TestIsString 测试字符串类型检查
func TestIsString(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected string
		isString bool
	}{
		{"hello", "hello", true},
		{"", "", true},
		{123, "", false},
		{12.34, "", false},
		{true, "", false},
		{nil, "", false},
		{[]string{"a", "b"}, "", false},
		{map[string]string{"key": "value"}, "", false},
	}

	for _, tc := range testCases {
		result, ok := IsString(tc.input)

		if ok != tc.isString {
			t.Errorf("输入 %v: 期望 isString=%v, 得到 %v", tc.input, tc.isString, ok)
		}

		if ok && result != tc.expected {
			t.Errorf("输入 %v: 期望结果 '%s', 得到 '%s'", tc.input, tc.expected, result)
		}
	}
}

// TestIsString_EdgeCases 测试边界情况
func TestIsString_EdgeCases(t *testing.T) {
	// 测试中文字符串
	str, ok := IsString("中文测试")
	if !ok || str != "中文测试" {
		t.Error("中文字符串检查失败")
	}

	// 测试特殊字符
	str, ok = IsString("hello\nworld\t!")
	if !ok || str != "hello\nworld\t!" {
		t.Error("特殊字符字符串检查失败")
	}

	// 测试空格
	str, ok = IsString("   ")
	if !ok || str != "   " {
		t.Error("空格字符串检查失败")
	}
}

// TestIsUint 测试uint类型检查
func TestIsUint(t *testing.T) {
	testCases := []struct {
		input  interface{}
		result uint
		isUint bool
	}{
		{uint(123), 123, true},
		{uint(0), 0, true},
		{uint(18446744073709551615), 18446744073709551615, true}, // max uint
		{int(123), 0, false},
		{int32(123), 0, false},
		{int64(123), 0, false},
		{float64(123), 0, false},
		{"123", 0, false},
		{nil, 0, false},
	}

	for _, tc := range testCases {
		result, ok := IsUint(tc.input)

		if ok != tc.isUint {
			t.Errorf("输入 %v: 期望 isUint=%v, 得到 %v", tc.input, tc.isUint, ok)
		}

		if ok && result != tc.result {
			t.Errorf("输入 %v: 期望结果 %d, 得到 %d", tc.input, tc.result, result)
		}
	}
}

// TestIsUint_EdgeCases 测试uint边界情况
func TestIsUint_EdgeCases(t *testing.T) {
	// 测试最大值
	maxUint := uint(18446744073709551615)
	result, ok := IsUint(maxUint)
	if !ok || result != maxUint {
		t.Error("uint最大值检查失败")
	}

	// 测试零值
	result, ok = IsUint(uint(0))
	if !ok || result != 0 {
		t.Error("uint零值检查失败")
	}
}

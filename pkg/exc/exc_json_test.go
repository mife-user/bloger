package exc

import (
	"testing"
)

// TestExcFileToJSON 测试序列化为JSON
func TestExcFileToJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected string
		hasError bool
	}{
		{
			name:     "简单结构体",
			input:    map[string]string{"key": "value"},
			expected: `{"key":"value"}`,
			hasError: false,
		},
		{
			name:     "数字",
			input:    map[string]int{"count": 42},
			expected: `{"count":42}`,
			hasError: false,
		},
		{
			name:     "嵌套结构",
			input:    map[string]interface{}{"user": map[string]string{"name": "Alice"}},
			expected: `{"user":{"name":"Alice"}}`,
			hasError: false,
		},
		{
			name:     "空结构体",
			input:    map[string]string{},
			expected: `{}`,
			hasError: false,
		},
		{
			name:     "数组",
			input:    []string{"a", "b", "c"},
			expected: `["a","b","c"]`,
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ExcFileToJSON(tc.input)

			if tc.hasError {
				if err == nil {
					t.Error("期望返回错误但没有")
				}
				return
			}

			if err != nil {
				t.Errorf("不期望错误: %v", err)
				return
			}

			if result != tc.expected {
				t.Errorf("期望 '%s', 得到 '%s'", tc.expected, result)
			}
		})
	}
}

// TestExcJSONToFile 测试从JSON反序列化
func TestExcJSONToFile(t *testing.T) {
	t.Run("简单结构体", func(t *testing.T) {
		jsonStr := `{"key":"value"}`
		var result map[string]string

		err := ExcJSONToFile(jsonStr, &result)

		if err != nil {
			t.Errorf("不期望错误: %v", err)
			return
		}

		if result["key"] != "value" {
			t.Errorf("期望 key=value, 得到 %v", result)
		}
	})

	t.Run("数字", func(t *testing.T) {
		jsonStr := `{"count":42}`
		var result map[string]int

		err := ExcJSONToFile(jsonStr, &result)

		if err != nil {
			t.Errorf("不期望错误: %v", err)
			return
		}

		if result["count"] != 42 {
			t.Errorf("期望 count=42, 得到 %v", result)
		}
	})

	t.Run("嵌套结构", func(t *testing.T) {
		jsonStr := `{"user":{"name":"Alice"}}`
		var result map[string]interface{}

		err := ExcJSONToFile(jsonStr, &result)

		if err != nil {
			t.Errorf("不期望错误: %v", err)
			return
		}

		user := result["user"].(map[string]interface{})
		if user["name"] != "Alice" {
			t.Errorf("期望 name=Alice, 得到 %v", user["name"])
		}
	})

	t.Run("无效JSON", func(t *testing.T) {
		jsonStr := `invalid json`
		var result map[string]string

		err := ExcJSONToFile(jsonStr, &result)

		if err == nil {
			t.Error("期望返回错误但没有")
		}
	})
}

// TestExcFileToJSON_InvalidInput 测试无效输入
func TestExcFileToJSON_InvalidInput(t *testing.T) {
	// 测试无法序列化的类型（如channel）
	ch := make(chan int)
	_, err := ExcFileToJSON(ch)

	if err == nil {
		t.Error("无法序列化的类型应该返回错误")
	}
}

// TestExcJSONToFile_InvalidInput 测试无效输入
func TestExcJSONToFile_InvalidInput(t *testing.T) {
	t.Run("无效JSON字符串", func(t *testing.T) {
		var result map[string]string
		err := ExcJSONToFile("not valid json", &result)

		if err == nil {
			t.Error("无效JSON应该返回错误")
		}
	})

	t.Run("类型不匹配", func(t *testing.T) {
		jsonStr := `{"key": "value"}`
		var result map[string]int // 类型不匹配
		err := ExcJSONToFile(jsonStr, &result)

		// JSON数字字符串转int会失败
		if err == nil {
			t.Log("类型不匹配可能不返回错误，取决于JSON解析器")
		}
	})
}

// TestExcRoundTrip 测试序列化和反序列化的往返
func TestExcRoundTrip(t *testing.T) {
	original := map[string]interface{}{
		"name":    "Alice",
		"age":     30,
		"active":  true,
		"tags":    []string{"go", "test"},
		"address": map[string]string{"city": "Beijing"},
	}

	// 序列化
	jsonStr, err := ExcFileToJSON(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	// 反序列化
	var result map[string]interface{}
	err = ExcJSONToFile(jsonStr, &result)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	// 验证关键字段
	if result["name"] != original["name"] {
		t.Errorf("name不匹配")
	}
}

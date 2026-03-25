package gitmodel

import (
	"encoding/json"
	"testing"
)

// TestGitModel_JSONMarshal 测试JSON序列化
func TestGitModel_JSONMarshal(t *testing.T) {
	model := GitModel{
		Token: "ghp_testtoken123456",
	}

	// 序列化
	jsonData, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	// 验证JSON内容
	expected := `{"token":"ghp_testtoken123456"}`
	if string(jsonData) != expected {
		t.Errorf("期望 %s, 得到 %s", expected, string(jsonData))
	}
}

// TestGitModel_JSONUnmarshal 测试JSON反序列化
func TestGitModel_JSONUnmarshal(t *testing.T) {
	jsonStr := `{"token":"ghp_testtoken123456"}`

	var model GitModel
	err := json.Unmarshal([]byte(jsonStr), &model)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	// 验证字段
	if model.Token != "ghp_testtoken123456" {
		t.Errorf("期望 token=ghp_testtoken123456, 得到 %s", model.Token)
	}
}

// TestGitModel_EmptyToken 测试空Token
func TestGitModel_EmptyToken(t *testing.T) {
	model := GitModel{}

	// 序列化空模型
	jsonData, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("序列化空模型失败: %v", err)
	}

	// 反序列化
	var loaded GitModel
	err = json.Unmarshal(jsonData, &loaded)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if loaded.Token != "" {
		t.Errorf("期望空Token, 得到 %s", loaded.Token)
	}
}

// TestGitModel_SpecialCharacters 测试特殊字符
func TestGitModel_SpecialCharacters(t *testing.T) {
	specialTokens := []string{
		"ghp_token!@#$%^&*()",
		"ghp_token with spaces",
		"ghp_token\nwith\nnewlines",
		"ghp_token\"with\"quotes",
	}

	for _, token := range specialTokens {
		model := GitModel{Token: token}

		// 序列化
		jsonData, err := json.Marshal(model)
		if err != nil {
			t.Errorf("序列化特殊字符Token失败: %v", err)
			continue
		}

		// 反序列化
		var loaded GitModel
		err = json.Unmarshal(jsonData, &loaded)
		if err != nil {
			t.Errorf("反序列化失败: %v", err)
			continue
		}

		// 验证Token一致
		if loaded.Token != token {
			t.Errorf("Token不匹配: 期望 %s, 得到 %s", token, loaded.Token)
		}
	}
}

// TestGitModel_LongToken 测试长Token
func TestGitModel_LongToken(t *testing.T) {
	// 生成一个很长的Token
	longToken := ""
	for i := 0; i < 1000; i++ {
		longToken += "a"
	}

	model := GitModel{Token: longToken}

	// 序列化
	jsonData, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("序列化长Token失败: %v", err)
	}

	// 反序列化
	var loaded GitModel
	err = json.Unmarshal(jsonData, &loaded)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if loaded.Token != longToken {
		t.Error("长Token不匹配")
	}
}

// TestGitModel_UnicodeToken 测试Unicode Token
func TestGitModel_UnicodeToken(t *testing.T) {
	unicodeTokens := []string{
		"ghp_中文token",
		"ghp_🔐🔒🔓",
		"ghp_emoji🎉token",
	}

	for _, token := range unicodeTokens {
		model := GitModel{Token: token}

		// 序列化
		jsonData, err := json.Marshal(model)
		if err != nil {
			t.Errorf("序列化Unicode Token失败: %v", err)
			continue
		}

		// 反序列化
		var loaded GitModel
		err = json.Unmarshal(jsonData, &loaded)
		if err != nil {
			t.Errorf("反序列化失败: %v", err)
			continue
		}

		if loaded.Token != token {
			t.Errorf("Unicode Token不匹配: 期望 %s, 得到 %s", token, loaded.Token)
		}
	}
}

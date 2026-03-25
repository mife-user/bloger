package gitrepo

import (
	"bloger/pkg/conf"
	"bloger/pkg/db"
	"os"
	"path/filepath"
	"testing"
)

// TestGitRepo_Save 测试保存Token
func TestGitRepo_Save(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	// 创建配置
	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	// 创建GitRepo
	repo := NewGitRepo(config)

	// 保存Token
	testToken := "ghp_testtoken123456"
	err := repo.Save(testToken)
	if err != nil {
		t.Fatalf("保存Token失败: %v", err)
	}

	// 验证文件是否存在
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatal("文件应该存在但不存在")
	}

	// 验证文件内容
	jsonDB := db.NewJSONFileDB(testFile)
	var loaded map[string]interface{}
	err = jsonDB.Load(&loaded)
	if err != nil {
		t.Fatalf("加载文件失败: %v", err)
	}

	// 验证Token字段存在
	if _, ok := loaded["token"]; !ok {
		t.Fatal("文件中应该包含token字段")
	}
}

// TestGitRepo_Save_Overwrite 测试覆盖保存
func TestGitRepo_Save_Overwrite(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	repo := NewGitRepo(config)

	// 第一次保存
	err := repo.Save("first_token")
	if err != nil {
		t.Fatalf("第一次保存失败: %v", err)
	}

	// 第二次保存（覆盖）
	err = repo.Save("second_token")
	if err != nil {
		t.Fatalf("第二次保存失败: %v", err)
	}

	// 验证是第二次的Token
	jsonDB := db.NewJSONFileDB(testFile)
	var loaded map[string]interface{}
	err = jsonDB.Load(&loaded)
	if err != nil {
		t.Fatalf("加载文件失败: %v", err)
	}

	if loaded["token"] != "second_token" {
		t.Errorf("期望 token=second_token, 得到 %v", loaded["token"])
	}
}

// TestGitRepo_Save_NestedDirectory 测试保存到嵌套目录
func TestGitRepo_Save_NestedDirectory(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "nested", "deep", "dir", "secrets.json")

	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	repo := NewGitRepo(config)

	// 保存到嵌套目录
	err := repo.Save("nested_token")
	if err != nil {
		t.Fatalf("保存到嵌套目录失败: %v", err)
	}

	// 验证文件存在
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatal("文件应该存在但不存在")
	}
}

// TestGitRepo_Save_EmptyToken 测试保存空Token
func TestGitRepo_Save_EmptyToken(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	repo := NewGitRepo(config)

	// 保存空Token
	err := repo.Save("")
	if err != nil {
		t.Fatalf("保存空Token失败: %v", err)
	}

	// 验证文件内容
	jsonDB := db.NewJSONFileDB(testFile)
	var loaded map[string]interface{}
	err = jsonDB.Load(&loaded)
	if err != nil {
		t.Fatalf("加载文件失败: %v", err)
	}

	if loaded["token"] != "" {
		t.Errorf("期望空token, 得到 %v", loaded["token"])
	}
}

// TestGitRepo_NewGitRepo 测试创建GitRepo
func TestGitRepo_NewGitRepo(t *testing.T) {
	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: "/test/path/secrets.json",
		},
	}

	repo := NewGitRepo(config)

	if repo == nil {
		t.Fatal("GitRepo不应该为nil")
	}

	if repo.db == nil {
		t.Fatal("GitRepo.db不应该为nil")
	}
}

// TestGitRepo_Save_SpecialCharacters 测试保存特殊字符Token
func TestGitRepo_Save_SpecialCharacters(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	repo := NewGitRepo(config)

	specialTokens := []string{
		"ghp_token!@#$%^&*()",
		"ghp_token with spaces",
		"ghp_token\nwith\nnewlines",
	}

	for _, token := range specialTokens {
		err := repo.Save(token)
		if err != nil {
			t.Errorf("保存特殊字符Token失败: %v", err)
			continue
		}

		// 验证内容
		jsonDB := db.NewJSONFileDB(testFile)
		var loaded map[string]interface{}
		err = jsonDB.Load(&loaded)
		if err != nil {
			t.Errorf("加载文件失败: %v", err)
			continue
		}

		if loaded["token"] != token {
			t.Errorf("Token不匹配: 期望 %s, 得到 %v", token, loaded["token"])
		}
	}
}

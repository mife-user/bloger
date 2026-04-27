package gitservice

import (
	"mifer/internal/repo/gitrepo"
	"mifer/pkg/conf"
	"mifer/pkg/db"
	"mifer/pkg/utils"
	"context"
	"os"
	"path/filepath"
	"testing"
)

// TestGitService_Save 测试服务层保存Token
func TestGitService_Save(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	// 创建配置
	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	// 创建依赖
	repo := gitrepo.NewGitRepo(config)
	service := NewGitService(repo)

	// 保存Token
	testToken := "ghp_testtoken123456"
	err := service.Save(context.Background(), testToken)
	if err != nil {
		t.Fatalf("保存Token失败: %v", err)
	}

	// 验证文件存在
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatal("文件应该存在但不存在")
	}

	// 验证Token被加密
	jsonDB := db.NewJSONFileDB(testFile)
	var loaded map[string]interface{}
	err = jsonDB.Load(&loaded)
	if err != nil {
		t.Fatalf("加载文件失败: %v", err)
	}

	// Token应该被加密，不应该是明文
	savedToken, ok := loaded["token"].(string)
	if !ok {
		t.Fatal("token字段应该是字符串")
	}

	// 验证保存的不是明文Token
	if savedToken == testToken {
		t.Fatal("Token应该被加密，不应该是明文")
	}

	// 验证可以验证密码
	if !utils.CheckPasswordHash(testToken, savedToken) {
		t.Fatal("加密后的Token应该能被验证")
	}
}

// TestGitService_Save_EmptyToken 测试保存空Token
func TestGitService_Save_EmptyToken(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	repo := gitrepo.NewGitRepo(config)
	service := NewGitService(repo)

	// 保存空Token
	err := service.Save(context.Background(), "")
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

	savedToken, ok := loaded["token"].(string)
	if !ok {
		t.Fatal("token字段应该是字符串")
	}

	// 空Token也应该被加密
	if savedToken == "" {
		t.Fatal("空Token也应该被加密")
	}

	// 验证空Token
	if !utils.CheckPasswordHash("", savedToken) {
		t.Fatal("加密后的空Token应该能被验证")
	}
}

// TestGitService_Save_DifferentTokens 测试保存不同Token
func TestGitService_Save_DifferentTokens(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	repo := gitrepo.NewGitRepo(config)
	service := NewGitService(repo)

	tokens := []string{
		"token1",
		"token2",
		"token_with_special_chars!@#$%",
		"token!@#$%^&*()",
	}

	for _, token := range tokens {
		// 保存Token
		err := service.Save(context.Background(), token)
		if err != nil {
			t.Errorf("保存Token失败: %v", err)
			continue
		}

		// 验证Token
		jsonDB := db.NewJSONFileDB(testFile)
		var loaded map[string]interface{}
		err = jsonDB.Load(&loaded)
		if err != nil {
			t.Errorf("加载文件失败: %v", err)
			continue
		}

		savedToken, ok := loaded["token"].(string)
		if !ok {
			t.Errorf("token字段应该是字符串")
			continue
		}

		// 验证Token
		if !utils.CheckPasswordHash(token, savedToken) {
			t.Errorf("Token验证失败")
		}
	}
}

// TestGitService_NewGitService 测试创建GitService
func TestGitService_NewGitService(t *testing.T) {
	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: "/test/path/secrets.json",
		},
	}

	repo := gitrepo.NewGitRepo(config)
	service := NewGitService(repo)

	if service == nil {
		t.Fatal("GitService不应该为nil")
	}

	if service.Repo == nil {
		t.Fatal("GitService.Repo不应该为nil")
	}
}

// TestGitService_Save_Overwrite 测试覆盖保存
func TestGitService_Save_Overwrite(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	config := &conf.Config{
		Git: conf.GitConfig{
			LockPath: testFile,
		},
	}

	repo := gitrepo.NewGitRepo(config)
	service := NewGitService(repo)

	// 第一次保存
	firstToken := "first_token"
	err := service.Save(context.Background(), firstToken)
	if err != nil {
		t.Fatalf("第一次保存失败: %v", err)
	}

	// 获取第一次的哈希
	jsonDB := db.NewJSONFileDB(testFile)
	var loaded1 map[string]interface{}
	jsonDB.Load(&loaded1)
	firstHash := loaded1["token"].(string)

	// 第二次保存
	secondToken := "second_token"
	err = service.Save(context.Background(), secondToken)
	if err != nil {
		t.Fatalf("第二次保存失败: %v", err)
	}

	// 获取第二次的哈希
	var loaded2 map[string]interface{}
	jsonDB.Load(&loaded2)
	secondHash := loaded2["token"].(string)

	// 两个哈希应该不同（bcrypt每次生成不同的哈希）
	if firstHash == secondHash {
		t.Fatal("两次保存的哈希应该不同")
	}

	// 验证第二次的Token
	if !utils.CheckPasswordHash(secondToken, secondHash) {
		t.Fatal("第二次的Token验证失败")
	}

	// 第一次的Token不应该再有效
	if utils.CheckPasswordHash(firstToken, secondHash) {
		t.Fatal("第一次的Token不应该再有效")
	}
}

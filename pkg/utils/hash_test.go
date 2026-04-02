package utils

import (
	"testing"
)

// TestHashPassword 测试密码哈希
func TestHashPassword(t *testing.T) {
	password := "mySecretPassword123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}

	// 哈希不应该为空
	if hash == "" {
		t.Fatal("哈希值不应该为空")
	}

	// 哈希值不应该等于原密码
	if hash == password {
		t.Fatal("哈希值不应该等于原密码")
	}

	// 哈希值长度应该大于原密码
	if len(hash) <= len(password) {
		t.Fatal("哈希值长度应该大于原密码")
	}
}

// TestHashPassword_DifferentHashes 测试相同密码生成不同哈希
func TestHashPassword_DifferentHashes(t *testing.T) {
	password := "mySecretPassword123"

	hash1, err := HashPassword(password)
	if err != nil {
		t.Fatalf("第一次哈希失败: %v", err)
	}

	hash2, err := HashPassword(password)
	if err != nil {
		t.Fatalf("第二次哈希失败: %v", err)
	}

	// bcrypt每次生成的哈希值应该不同（因为有随机盐）
	if hash1 == hash2 {
		t.Fatal("相同密码的两次哈希值应该不同")
	}
}

// TestHashPassword_EmptyPassword 测试空密码
func TestHashPassword_EmptyPassword(t *testing.T) {
	password := ""

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("哈希空密码失败: %v", err)
	}

	// 空密码也应该能被哈希
	if hash == "" {
		t.Fatal("空密码的哈希值不应该为空")
	}
}

// TestCheckPasswordHash 测试密码验证
func TestCheckPasswordHash(t *testing.T) {
	password := "mySecretPassword123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}

	// 正确的密码应该验证通过
	if !CheckPasswordHash(password, hash) {
		t.Fatal("正确密码验证失败")
	}

	// 错误的密码应该验证失败
	if CheckPasswordHash("wrongPassword", hash) {
		t.Fatal("错误密码不应该验证通过")
	}
}

// TestCheckPasswordHash_DifferentPasswords 测试不同密码的哈希验证
func TestCheckPasswordHash_DifferentPasswords(t *testing.T) {
	passwords := []string{
		"simple",
		"with spaces",
		"with!special@chars#",
		"数字密码123",
		"verylongpasswordverylongpasswordverylongpasswordverylongpassword",
	}

	for _, password := range passwords {
		hash, err := HashPassword(password)
		if err != nil {
			t.Errorf("哈希密码 '%s' 失败: %v", password, err)
			continue
		}

		if !CheckPasswordHash(password, hash) {
			t.Errorf("密码 '%s' 验证失败", password)
		}
	}
}

// TestCheckPasswordHash_EmptyPassword 测试空密码验证
func TestCheckPasswordHash_EmptyPassword(t *testing.T) {
	password := ""

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("哈希空密码失败: %v", err)
	}

	// 空密码也应该能验证通过
	if !CheckPasswordHash(password, hash) {
		t.Fatal("空密码验证失败")
	}

	// 非空密码不应该验证通过
	if CheckPasswordHash("notempty", hash) {
		t.Fatal("非空密码不应该验证通过空密码的哈希")
	}
}

// TestCheckPasswordHash_InvalidHash 测试无效哈希
func TestCheckPasswordHash_InvalidHash(t *testing.T) {
	password := "myPassword"

	// 无效的哈希格式
	invalidHashes := []string{
		"",
		"notavalidhash",
		"tooshort",
	}

	for _, invalidHash := range invalidHashes {
		if CheckPasswordHash(password, invalidHash) {
			t.Errorf("无效哈希 '%s' 不应该验证通过", invalidHash)
		}
	}
}

// TestHashPassword_LongPassword 测试长密码
func TestHashPassword_LongPassword(t *testing.T) {
	// bcrypt有72字节的限制
	longPassword := ""
	for i := 0; i < 100; i++ {
		longPassword += "a"
	}

	hash, err := HashPassword(longPassword)

	// bcrypt会返回错误当密码超过72字节
	if err != nil {
		// 这是预期行为，bcrypt有72字节限制
		t.Logf("长密码超过bcrypt限制（72字节）: %v", err)
		return
	}

	// 如果没有错误，验证密码
	if !CheckPasswordHash(longPassword, hash) {
		t.Fatal("长密码验证失败")
	}
}

// TestHashPassword_Unicode 测试Unicode密码
func TestHashPassword_Unicode(t *testing.T) {
	passwords := []string{
		"密码测试",
		"🔐🔒🔓",
		"混合password123中文",
	}

	for _, password := range passwords {
		hash, err := HashPassword(password)
		if err != nil {
			t.Errorf("哈希Unicode密码失败: %v", err)
			continue
		}

		if !CheckPasswordHash(password, hash) {
			t.Errorf("Unicode密码 '%s' 验证失败", password)
		}
	}
}

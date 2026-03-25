package db

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestJSONFileDB_Save 测试保存功能
func TestJSONFileDB_Save(t *testing.T) {
	// 创建临时测试目录
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.json")

	db := NewJSONFileDB(testFile)

	// 测试数据
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	// 保存数据
	err := db.Save(data)
	if err != nil {
		t.Fatalf("保存数据失败: %v", err)
	}

	// 验证文件是否存在
	if !db.Exists() {
		t.Fatal("文件应该存在但不存在")
	}

	// 读取文件内容验证
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("打开文件失败: %v", err)
	}
	defer file.Close()

	var loadedData map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&loadedData); err != nil {
		t.Fatalf("解码JSON失败: %v", err)
	}

	// 验证数据内容
	if loadedData["key1"] != "value1" {
		t.Errorf("期望 key1=value1, 得到 %s", loadedData["key1"])
	}
	if loadedData["key2"] != "value2" {
		t.Errorf("期望 key2=value2, 得到 %s", loadedData["key2"])
	}
}

// TestJSONFileDB_Save_NestedDirectory 测试保存到嵌套目录
func TestJSONFileDB_Save_NestedDirectory(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "nested", "dir", "test.json")

	db := NewJSONFileDB(testFile)

	data := map[string]int{"count": 42}

	// 保存数据到嵌套目录
	err := db.Save(data)
	if err != nil {
		t.Fatalf("保存到嵌套目录失败: %v", err)
	}

	// 验证文件存在
	if !db.Exists() {
		t.Fatal("文件应该存在但不存在")
	}
}

// TestJSONFileDB_Load 测试加载功能
func TestJSONFileDB_Load(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.json")

	// 准备测试数据
	originalData := map[string]interface{}{
		"name":  "test",
		"value": 123,
	}

	// 写入文件
	jsonData, _ := json.MarshalIndent(originalData, "", "  ")
	if err := os.WriteFile(testFile, jsonData, 0644); err != nil {
		t.Fatalf("写入测试文件失败: %v", err)
	}

	db := NewJSONFileDB(testFile)

	// 加载数据
	var loadedData map[string]interface{}
	err := db.Load(&loadedData)
	if err != nil {
		t.Fatalf("加载数据失败: %v", err)
	}

	// 验证数据
	if loadedData["name"] != "test" {
		t.Errorf("期望 name=test, 得到 %v", loadedData["name"])
	}
	if int(loadedData["value"].(float64)) != 123 {
		t.Errorf("期望 value=123, 得到 %v", loadedData["value"])
	}
}

// TestJSONFileDB_Load_FileNotExists 测试加载不存在的文件
func TestJSONFileDB_Load_FileNotExists(t *testing.T) {
	db := NewJSONFileDB("/nonexistent/path/file.json")

	var data map[string]string
	err := db.Load(&data)

	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
}

// TestJSONFileDB_Exists 测试文件存在检查
func TestJSONFileDB_Exists(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.json")

	db := NewJSONFileDB(testFile)

	// 文件不存在
	if db.Exists() {
		t.Fatal("文件不应该存在")
	}

	// 创建文件
	if err := os.WriteFile(testFile, []byte("{}"), 0644); err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	// 文件存在
	if !db.Exists() {
		t.Fatal("文件应该存在")
	}
}

// TestJSONFileDB_Delete 测试删除功能
func TestJSONFileDB_Delete(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.json")

	db := NewJSONFileDB(testFile)

	// 创建文件
	data := map[string]string{"test": "data"}
	if err := db.Save(data); err != nil {
		t.Fatalf("保存测试文件失败: %v", err)
	}

	// 验证文件存在
	if !db.Exists() {
		t.Fatal("文件应该存在")
	}

	// 删除文件
	err := db.Delete()
	if err != nil {
		t.Fatalf("删除文件失败: %v", err)
	}

	// 验证文件不存在
	if db.Exists() {
		t.Fatal("文件不应该存在")
	}
}

// TestJSONFileDB_Delete_FileNotExists 测试删除不存在的文件
func TestJSONFileDB_Delete_FileNotExists(t *testing.T) {
	db := NewJSONFileDB("/nonexistent/path/file.json")

	err := db.Delete()

	if err == nil {
		t.Fatal("期望返回错误但没有")
	}
}

// TestJSONFileDB_SaveAndLoad 测试保存和加载的完整流程
func TestJSONFileDB_SaveAndLoad(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.json")

	db := NewJSONFileDB(testFile)

	// 定义测试结构体
	type TestStruct struct {
		Name    string
		Age     int
		Active  bool
		Tags    []string
	}

	original := TestStruct{
		Name:   "Alice",
		Age:    30,
		Active: true,
		Tags:   []string{"go", "test", "db"},
	}

	// 保存
	if err := db.Save(original); err != nil {
		t.Fatalf("保存失败: %v", err)
	}

	// 加载
	var loaded TestStruct
	if err := db.Load(&loaded); err != nil {
		t.Fatalf("加载失败: %v", err)
	}

	// 验证所有字段
	if loaded.Name != original.Name {
		t.Errorf("Name: 期望 %s, 得到 %s", original.Name, loaded.Name)
	}
	if loaded.Age != original.Age {
		t.Errorf("Age: 期望 %d, 得到 %d", original.Age, loaded.Age)
	}
	if loaded.Active != original.Active {
		t.Errorf("Active: 期望 %v, 得到 %v", original.Active, loaded.Active)
	}
	if len(loaded.Tags) != len(original.Tags) {
		t.Errorf("Tags长度: 期望 %d, 得到 %d", len(original.Tags), len(loaded.Tags))
	}
}

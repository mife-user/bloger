package db

import (
	"mifer/pkg/logger"
	"encoding/json"
	"os"
	"path/filepath"
)

// JSONFileDB JSON文件数据库
type JSONFileDB struct {
	filePath string
}

// NewJSONFileDB 新建JSON文件数据库
func NewJSONFileDB(filePath string) *JSONFileDB {
	return &JSONFileDB{
		filePath: filePath,
	}
}

// Save 保存文件
func (db *JSONFileDB) Save(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		logger.Error("序列化数据失败", logger.C(err))
		return err
	}

	dir := filepath.Dir(db.filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			logger.Error("创建目录失败", logger.C(err))
			return err
		}
	}

	file, err := os.OpenFile(db.filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logger.Error("打开文件失败", logger.C(err))
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		logger.Error("写入文件失败", logger.C(err))
		return err
	}

	logger.Info("数据保存成功", logger.S("path", db.filePath))
	return nil
}

// Load 载入文件
func (db *JSONFileDB) Load(data interface{}) error {
	file, err := os.Open(db.filePath)
	if err != nil {
		logger.Error("打开文件失败", logger.C(err))
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		logger.Error("反序列化数据失败", logger.C(err))
		return err
	}

	logger.Info("数据加载成功", logger.S("path", db.filePath))
	return nil
}

// Exists 检查文件是否存在
func (db *JSONFileDB) Exists() bool {
	_, err := os.Stat(db.filePath)
	return !os.IsNotExist(err)
}

// Delete 删除文件
func (db *JSONFileDB) Delete() error {
	if err := os.Remove(db.filePath); err != nil {
		logger.Error("删除文件失败", logger.C(err))
		return err
	}
	logger.Info("文件删除成功", logger.S("path", db.filePath))
	return nil
}

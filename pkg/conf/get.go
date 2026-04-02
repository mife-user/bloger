package conf

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *Config

func GetConfig() *Config {
	return config
}

// LoadConfig 加载配置文件
func LoadConfig() error {
	// 加载配置文件
	v := viper.New()

	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// 尝试多个可能的配置文件路径
	configPaths := []string{
		filepath.Join(workDir, "config"),
		filepath.Join(workDir, "../../config"),
		"./config",
		"../../config",
	}

	found := false
	for _, path := range configPaths {
		if _, err := os.Stat(path); err == nil {
			v.AddConfigPath(path)
			found = true
			break
		}
	}

	if !found {
		// 使用默认路径
		v.AddConfigPath("../../config")
	}

	v.SetConfigName("dev")
	v.SetConfigType("yml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// 解析配置文件
	if err := v.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

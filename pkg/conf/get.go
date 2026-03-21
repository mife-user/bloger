package conf

import "github.com/spf13/viper"

var config *Config

func GetConfig() *Config {
	return config
}

// LoadConfig 加载配置文件
func LoadConfig() error {
	// 加载配置文件
	v := viper.New()
	// 加载配置文件路径
	path := "../../config"
	v.AddConfigPath(path)
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

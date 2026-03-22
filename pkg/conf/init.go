package conf

type Config struct {
	Env string    `mapstructure:"env"`
	Gin GinConfig `mapstructure:"gin"`
	Git GitConfig `mapstructure:"git"`
	Ai  AiConfig  `mapstructure:"ai"`
}

type AiConfig struct {
	ApiKey  string `mapstructure:"api_key"`
	BaseURL string `mapstructure:"base_url"`
	Model   string `mapstructure:"model"`
}

type GitConfig struct {
	LockPath string `mapstructure:"lock_path"`
}

type GinConfig struct {
	Mode string     `mapstructure:"mode"`
	Port string     `mapstructure:"port"`
	Cors CorsConfig `mapstructure:"cors"`
}

type CorsConfig struct {
	AllowOrigins []string `mapstructure:"allow_origins"`
	AllowMethods []string `mapstructure:"allow_methods"`
}

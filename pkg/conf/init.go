package conf

type Config struct {
	Env string    `mapstructure:"env"`
	Gin GinConfig `mapstructure:"gin"`
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

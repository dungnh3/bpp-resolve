package config

type Config struct {
	conf.Base   `mapstructure:",squash"`
	PostgreSQL  database.PostgreSQLConfig `json:"postgresql" mapstructure:"postgresql" yaml:"postgresql"`
	GormLogMode string                    `json:"gorm_log_mode" mapstructure:"gorm_log_mode" yaml:"gorm_log_mode"`
	Redis       rediz.Config              `json:"redis" mapstructure:"redis" yaml:"redis"`
}

type Base struct {
	Env    string        `json:"env" mapstructure:"env" yaml:"env"`
	Server server.Config `json:"server" mapstructure:"server" yaml:"server"`
	Logger log.Config    `json:"log" mapstructure:"log" yaml:"log"`
}
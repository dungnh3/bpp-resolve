package config

import (
	"bytes"
	"encoding/json"
	"github.com/dungnh3/bpp-resolve/pkg/conf"
	"github.com/dungnh3/bpp-resolve/pkg/database"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	conf.Base `mapstructure:",squash"`
	MySQL     database.MySQLConfig `json:"mysql" mapstructure:"mysql" yaml:"mysql"`
}

// loadDefaultConfig return a default object configuration
func loadDefaultConfig() *Config {
	return &Config{
		Base:  *conf.DefaultBaseConfig(),
		MySQL: database.MySQLDefaultConfig(),
	}
}

func Load() (*Config, error) {
	// You should set default config value here
	c := loadDefaultConfig()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Read config file failed. ", err)

		configBuffer, err := json.Marshal(c)

		if err != nil {
			return nil, err
		}
		_ = viper.ReadConfig(bytes.NewBuffer(configBuffer))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	// -- end of hacking --//
	viper.AutomaticEnv()
	err = viper.Unmarshal(c)
	return c, err
}
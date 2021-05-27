package server

import "fmt"

type Config struct {
	HTTP Listen `json:"http" mapstructure:"http" yaml:"http"`
}

type Listen struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

func DefaultConfig() Config {
	return Config{
		HTTP: Listen{
			Host: "0.0.0.0",
			Port: 8080,
		},
	}
}

// String return socket listen data source name
func (l Listen) String() string {
	return fmt.Sprintf("%v:%v", l.Host, l.Port)
}

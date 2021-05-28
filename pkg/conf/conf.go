package conf

import (
	"github.com/dungnh3/bpp-resolve/pkg/log"
	"github.com/dungnh3/bpp-resolve/pkg/server"
)

// deploy env.
const (
	DeployEnvDev  = "dev"
	DeployEnvStag = "stag"
	DeployEnvProd = "prod"
)

type Base struct {
	Env    string        `json:"env" mapstructure:"env" yaml:"env"`
	Server server.Config `json:"server" mapstructure:"server" yaml:"server"`
	Logger log.Config    `json:"log" mapstructure:"log" yaml:"log"`
}

func DefaultBaseConfig() *Base {
	return &Base{
		Env:    DeployEnvDev,
		Server: server.DefaultConfig(),
		Logger: log.DefaultConfig(),
	}
}

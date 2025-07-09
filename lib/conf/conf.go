package conf

import (
	"github.com/creasty/defaults"
)

type Config struct {
	Language string            `yaml:"language" default:"zh-CN"`
	Urls     map[string]string `yaml:"urls"`
	Env      string            `yaml:"env" default:"dev"`

	Traefik *TraefikConfig `yaml:"traefik"`
}

type TraefikConfig struct {
	ConfPath string `yaml:"confPath" default:"conf/traefik"`
}

var Cfg *Config = nil

func Init(path string) {
	Cfg = &Config{}
	err := defaults.Set(Cfg)
	if err != nil {
		panic(err)
	}
}

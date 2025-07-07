package conf

import "trapolit/lib/i18n"

type Config struct {
	Language i18n.Lang
}

var Cfg *Config = nil

func Init(path string) {
	//TODO: 完成读取配置的相关代码
}

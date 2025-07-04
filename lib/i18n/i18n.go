package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"sync"
)

var bundle *i18n.Bundle
var once sync.Once

type Lang string

const (
	CN_ZH Lang = "zh-CN"
	EN_US      = "en-US"
)

func Init() {
	once.Do(func() {
		bundle = i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

		//加载英语i18n配置
		_, err := bundle.LoadMessageFile("/conf/i18n/en.yaml")
		if err != nil {
			panic(err)
		}
		//加载简体中文i18n配置
		_, err = bundle.LoadMessageFile("/conf/i18n/cn.yaml")
		if err != nil {
			panic(err)
		}
	})
}

func Translate(lang Lang, tag string) string {
	loc := i18n.NewLocalizer(bundle, string(lang))
	res, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID: tag,
	})
	if err != nil {
		panic(err)
	}
	return res
}

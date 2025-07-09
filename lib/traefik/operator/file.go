package operator

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	traconf "trapolit/lib/traefik/conf"
)

//TODO:这里应该能支持不同的实验的操作对于外部 接口应该只需要提供映射关系就行

type ProviderOption struct {
	file io.Writer // 对应配置文件的io

	FilePath  string // 配置文件的目录
	Directory string `yaml:"directory"`
	Watch     bool   `yaml:"watch"`
	Name      string `yaml:"name"`
}

// DynaOption 对于动态文件每一个服务应该对应一个文件 不同服务的路由配置放在不同的文件中
type DynaOption struct {
	file     *os.File
	closed   bool
	FilePath string // 配置文件目录

	NowConf *traconf.DynamicConf // 当前的配置
}

type FileOperator struct {
	Provider *ProviderOption
	Dynamic  *DynaOption
}

func (f FileOperator) AddProvider() {
	//TODO implement me
	panic("implement me")
}

func (f FileOperator) AddService(name string, targetUrl []string) {
	if f.Dynamic.closed {
		return
	}
	services := f.Dynamic.NowConf.Http.Services
	u, ok := services[name]
	urls := make([]*traconf.Url, len(targetUrl))
	for i, v := range targetUrl {
		urls[i] = &traconf.Url{Url: v}
	}
	if ok {
		u.Servers = urls
	} else {
		services[name] = &traconf.LoadBalancer{Servers: urls}
	}
}

func (f FileOperator) AddRouter(name, originUrl, serviceName string) {
	if f.Dynamic.closed {
		return
	}
	router := f.Dynamic.NowConf.Http.Routers
	r, ok := router[name]
	rule := fmt.Sprintf("Host(`%s`)", originUrl)
	entry := &traconf.EntryPoint{
		Web: &traconf.Addr{Address: ":80"},
	}
	if ok {
		r.Rule = rule
		r.Service = serviceName
		r.EntryPoints = entry
	} else {
		router[name] = &traconf.Router{
			Rule:        rule,
			Service:     serviceName,
			EntryPoints: entry,
		}
	}
}

func (f FileOperator) DelProvider(name string) {
	//TODO implement me
	panic("implement me")
}

func (f FileOperator) DelRouter(name string) {
	if f.Dynamic.closed {
		return
	}
}

func (f FileOperator) DelService(name string) {
	if f.Dynamic.closed {
		return
	}
}

func (f FileOperator) Flush() error {
	byt, err := yaml.Marshal(f.Dynamic.NowConf)
	if err != nil {
		return err
	}
	_, err = f.Dynamic.file.Write(byt)
	if err != nil {
		return err
	}
	f.Dynamic.closed = true
	return f.Dynamic.file.Close()
}

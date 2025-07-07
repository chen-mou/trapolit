package operator

import "io"

type ProviderOption struct {
	file io.Writer // 对应配置文件的io

	FilePath  string // 配置文件的目录
	Directory string `yaml:"directory"`
	Watch     bool   `yaml:"watch"`
	Name      string `yaml:"name"`
}

// DynaOption 对于动态文件每一个服务应该对应一个文件 不同服务的路由配置放在不同的文件中
type DynaOption struct {
	file     io.Writer
	FilePath string // 配置文件目录

	//服务相关
	ServiceName string
	ServiceUrl  string // 服务路由

	//路由相关
	RouterName string
	RouterRule string //匹配路径
}

type FileOperator struct {
	Provider *ProviderOption
	Dynamic  *DynaOption
}

func (f FileOperator) AddProvider() {
	//TODO implement me
	panic("implement me")
}

func (f FileOperator) AddService() {
	//TODO implement me
	panic("implement me")
}

func (f FileOperator) AddRouter() {
	//TODO implement me
	panic("implement me")
}

func (f FileOperator) DelProvider(name string) {
	//TODO implement me
	panic("implement me")
}

func (f FileOperator) DelRouter(name string) {
	//TODO implement me
	panic("implement me")
}

func (f FileOperator) DelService(name string) {
	//TODO implement me
	panic("implement me")
}

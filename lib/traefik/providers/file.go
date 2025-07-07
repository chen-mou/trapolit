package providers

type DynamicConfig struct {
	HTTP HTTPConfig `yaml:"http"`
}

type HTTPConfig struct {
	Routers     map[string]Router     `yaml:"routers"`
	Services    map[string]Service    `yaml:"services"`
	Middlewares map[string]Middleware `yaml:"middlewares"`
}

type Router struct {
	Rule        string   `yaml:"rule"`
	EntryPoints []string `yaml:"entryPoints"`
	Service     string   `yaml:"service"`
	Middlewares []string `yaml:"middlewares"`
}

type Service struct {
	LoadBalancer LoadBalancer `yaml:"loadBalancer"`
}

type LoadBalancer struct {
	Servers []Server `yaml:"servers"`
}

type Server struct {
	URL string `yaml:"url"`
}

type Middleware struct {
	AddPrefix *AddPrefix `yaml:"addPrefix,omitempty"`
}

type AddPrefix struct {
	Prefix string `yaml:"prefix"`
}

type StaticConfig struct {
	Providers Config `yaml:"providers"`
}

type Config struct {
	File FileProvider `yaml:"file"`
}

type FileProvider struct {
	Directory string `yaml:"directory"`
	Watch     bool   `yaml:"watch"`
}

func DefaultDynConf() *DynamicConfig {
	return &DynamicConfig{}
}

func DefaultStaticConf() *StaticConfig {

}

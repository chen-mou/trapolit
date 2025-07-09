package conf

type DynamicConf struct {
	Http *Http `yaml:"http"`
}

type Http struct {
	Routers  map[string]*Router       `yaml:"routers" validate:"required"`
	Services map[string]*LoadBalancer `yaml:"services" validate:"required"`
}

type WeightServices struct {
	Name   string `yaml:"name"`
	Weight int    `yaml:"weight"`
}

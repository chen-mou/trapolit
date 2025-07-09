package conf

type LoadBalancer struct {
	Servers []*Url `yaml:"servers" validate:"required"`
}

type Url struct {
	Url string `yaml:"url"`
}

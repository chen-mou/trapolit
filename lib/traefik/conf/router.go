package conf

type Router struct {
	Rule        string      `yaml:"rule" validate:"required"`
	Service     string      `yaml:"service" validate:"required"`
	EntryPoints *EntryPoint `yaml:"entryPoints,omitempty" validate:"required"`
}

type EntryPoint struct {
	Web       *Addr `yaml:"web,omitempty" validate:"required"`
	Websecure *Addr `yaml:"websecure,omitempty" validate:"required"`
}

type Addr struct {
	Address string `yaml:"address"`
}

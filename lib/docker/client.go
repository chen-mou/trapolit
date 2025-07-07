package docker

import (
	"errors"
	"github.com/docker/docker/client"
	"sync"
	"trapolit/lib/conf"
	"trapolit/lib/i18n"
)

type Options struct {
	Host string `yaml:"host"`
	Auth *Auth  `yaml:"auth"`
}

type Auth struct {
	CaCertPath     string `json:"ca_cert_path,omitempty" yaml:"ca-cert-path"`
	ClientCertPath string `json:"client_cert_path,omitempty" yaml:"client-cert-path"`
	ClientKeyPath  string `json:"client_key_path,omitempty" yaml:"client-key-path"`
}

var once sync.Once
var Client *client.Client

func Init(opt *Options) {
	var err error
	once.Do(func() {
		Client, err = NewClient(opt)
		if err != nil {
			panic(err)
		}
	})
}

func NewClient(opt *Options) (*client.Client, error) {
	cli, err := client.NewClientWithOpts(toOpts(opt)...)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func toOpts(opt *Options) []client.Opt {
	res := make([]client.Opt, 0)
	if opt.Auth != nil {
		auth := opt.Auth
		res = append(res, client.WithTLSClientConfig(auth.CaCertPath, auth.ClientCertPath, auth.ClientKeyPath))
	}
	if opt.Host == "" {
		panic(errors.New(i18n.Translate(conf.Cfg.Language, "ERROR.DOCKER.HOST_NOT_FOUND")))
	}
	res = append(res, client.WithHost(opt.Host), client.WithAPIVersionNegotiation())
	return res
}

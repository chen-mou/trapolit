package operator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	traconf "trapolit/lib/traefik/conf"
)

type Builder interface {
	NewOperator(path string) (Operator, error)
}

type FileOperatorBuilder struct {
}

func NewOperator(path string) (Operator, error) {
	file, err := os.Open(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
		file, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	}
	var byt []byte
	tmp := make([]byte, 1024)
	for _, err = file.Read(tmp); err == nil; {
		byt = append(byt, tmp...)
	}
	if err != io.EOF {
		return nil, err
	}
	opt := &traconf.DynamicConf{}
	err = yaml.Unmarshal(byt, opt)
	if err != nil {
		return nil, err
	}
	validate := validator.New()
	err = validate.Struct(opt)
	if err != nil {
		return nil, err
	}

	return &FileOperator{
		Dynamic: &DynaOption{
			file:     file,
			FilePath: path,
			NowConf:  opt,
		},
	}, nil
}

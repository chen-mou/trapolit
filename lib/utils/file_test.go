package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReadYaml(t *testing.T) {
	yaml, err := ReadYaml("conf/traefik/test.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println(json.Marshal(yaml))
}

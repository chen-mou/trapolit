package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"trapolit/internal/entity"
	"trapolit/lib/conf"
	"trapolit/lib/docker"
	"trapolit/lib/traefik/operator"
	"trapolit/lib/utils"
)

type TestParam struct {
	ImageName string
	Tag       string
	Env       entity.TicketEnv
}

// Test 对于各种实验流程的抽象 是每个流程的最小单位
type Test interface {
	Do(ctx context.Context, param *TestParam) error //执行
	Rollback(ctx context.Context)                   //回滚
}

type ABTest struct {
	Operator operator.Operator
}

// Do
// TODO: 这里应该还得把容器的url作为参数传入provider operator中
func (ab *ABTest) Do(ctx context.Context, param *TestParam) error {
	name := param.ImageName + "-" + param.Tag + "-" + utils.NowFormat("060102150405")
	port, err := ab.findImagePorts(ctx, param)
	if err != nil {
		return err
	}
	imageName := param.ImageName + ":" + param.Tag
	aId, bId, err := ab.funcName(ctx, imageName, port, name)
	if err != nil {
		return err
	}

	aContainer, err := docker.ListContainerById(ctx, aId)
	if err != nil {
		return err
	}
	bContainer, err := docker.ListContainerById(ctx, bId)
	if err != nil {
		return err
	}
	aUrl, err := getContainerUrl(aContainer)
	if err != nil {
		return err
	}
	bUrl, err := getContainerUrl(bContainer)
	if err != nil {
		return err
	}
	//TODO: 编写把路由写入配置的逻辑
}

func (ab *ABTest) funcName(ctx context.Context, imageName string, port nat.PortMap, name string) (string, string, error) {
	aId, err := docker.CreateContainer(ctx, &docker.CreateContainOpt{
		Ports:     port,
		Volumes:   nil,
		Env:       nil,
		ImageName: imageName,
		Name:      name + "-a",
	})
	if err != nil {
		return "", "", err
	}
	bId, err := docker.CreateContainer(ctx, &docker.CreateContainOpt{
		Ports:     port,
		Volumes:   nil,
		Env:       nil,
		ImageName: imageName,
		Name:      name + "-b",
	})
	if err != nil {
		return "", "", err
	}
	return aId, bId, nil
}

func (ab *ABTest) findImagePorts(ctx context.Context, param *TestParam) (nat.PortMap, error) {
	image, err := docker.ListImageByName(ctx, param.ImageName+":"+param.Tag)
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, utils.NewError(conf.Cfg.Language, "ERROR.DOCKER.IMAGE_NOT_FOUND")
	}

	port := nat.PortMap{}
	for k := range image.Config.ExposedPorts {
		port[nat.Port(k)] = []nat.PortBinding{
			{
				HostPort: "",
			},
		}
	}
	return port, nil
}

// TODO: 这里应该有个配置每个容器应该要有个默认的端口用于导出 这个默认端口应该是可配置的 暂时为80
func getContainerUrl(container *container.InspectResponse) (string, error) {
	ip := ""
	port := ""
	for _, network := range container.NetworkSettings.Networks {
		if network.IPAddress != "" {
			ip = network.IPAddress
			break
		}
	}
	for k, p := range container.HostConfig.PortBindings {
		if k == "80" && len(p) != 0 {
			port = p[0].HostPort
			break
		}
	}
	if ip == "" || port == "" {
		return "", utils.NewError(conf.Cfg.Language, "DOCKER.ERROR.INVALID_CONTAINER")
	}
	return ip + ":" + port, nil // 没找到
}

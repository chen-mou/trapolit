package docker

import (
	"context"
	"github.com/docker/docker/api/types/build"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"os"
	"time"
	"trapolit/lib/utils"
)

type CreateContainOpt struct {
	Ports     nat.PortMap
	Volumes   []mount.Mount
	Env       []string
	ImageName string
	Name      string
	Labels    map[string]string
}

// BuildImage 将dockerfile打包为镜像
func BuildImage(ctx context.Context, imageName string, path string) error {
	name := utils.MD5(imageName + time.Now().Format("20060102150405"))
	targetFile := "/" + name + ".tar"
	// 将目标文件夹打包成tar
	err := utils.TarDirectory(path, targetFile)
	if err != nil {
		return err
	}
	file, err := os.Open(targetFile)
	if err != nil {
		return err
	}
	_, err = Client.ImageBuild(ctx, file, build.ImageBuildOptions{})
	if err != nil {
		return err
	}
	return nil
}

// CreateContainer 创建容器
func CreateContainer(ctx context.Context, opt *CreateContainOpt) (string, error) {
	var hostConfig *container.HostConfig = nil
	config := &container.Config{Image: opt.ImageName}
	now := time.Now().Format("20060102")
	containerName := opt.ImageName + "_" + now
	if opt.Name != "" {
		containerName = opt.Name
	}
	if opt.Ports != nil {
		hostConfig = &container.HostConfig{}
		hostConfig.PortBindings = opt.Ports
	}
	if opt.Volumes != nil {
		if hostConfig == nil {
			hostConfig = &container.HostConfig{}
		}
		hostConfig.Mounts = opt.Volumes
	}
	if opt.Env != nil {
		config.Env = opt.Env
	}
	if opt.Labels != nil {
		config.Labels = opt.Labels
	}
	create, err := Client.ContainerCreate(ctx,
		config,
		hostConfig,
		&network.NetworkingConfig{},
		&v1.Platform{},
		containerName,
	)
	if err != nil {
		return "", err
	}
	return create.ID, nil
}

func ListContainerById(ctx context.Context, containerId string) (*container.InspectResponse, error) {
	inspect, err := Client.ContainerInspect(ctx, containerId)
	if err != nil {
		return nil, err
	}
	return &inspect, nil
}

func ListImageByName(ctx context.Context, name string) (*image.InspectResponse, error) {
	args := filters.Args{}
	args.Add("reference", name)
	list, err := Client.ImageList(ctx, image.ListOptions{Filters: args})
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	inspect, err := Client.ImageInspect(ctx, list[0].ID)
	if err != nil {
		return nil, err
	}
	return &inspect, nil
}

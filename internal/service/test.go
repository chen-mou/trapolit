package service

import (
	"context"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"time"
	"trapolit/internal/entity"
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

// Do TODO: 创建ab两个容器暴露端口 更具容器名与基础域名拼出源域名 配置中分别写入路由
func (ab *ABTest) Do(ctx context.Context, param *TestParam) error {
	name := param.ImageName + "-" + param.Tag + "-" + utils.NowFormat("060102150405")
	a, err := docker.CreateContainer(ctx, &docker.CreateContainOpt{
		Ports:     nil,
		Volumes:   nil,
		Env:       nil,
		ImageName: param.ImageName + ":" + param.Tag,
		Name:      name + "-a",
	})
	if err != nil {
		return err
	}
	b, err := docker.CreateContainer(ctx, &docker.CreateContainOpt{
		Ports:     nil,
		Volumes:   nil,
		Env:       nil,
		ImageName: param.ImageName + ":" + param.Tag,
		Name:      name + "-b",
	})
	if err != nil {
		return err
	}
}

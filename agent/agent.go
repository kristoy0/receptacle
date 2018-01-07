package main

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	unit "github.com/docker/go-units"
	proto "github.com/kristoy0/receptacle/server/proto"
	micro "github.com/micro/go-micro"
)

type TaskHandler interface {
	Deploy(context.Context, *proto.DeployRequest, *proto.DeployResponse) error
	Undeploy(context.Context, *proto.UndeployRequest, *proto.UndeployResponse) error
	List(context.Context, *proto.ListRequest, *proto.ListResponse) error
}

type Task struct{}

func (t *Task) Deploy(ctx context.Context, req *proto.DeployRequest, res *proto.DeployResponse) error {
	cli, err := getEnv()
	if err != nil {
		return err
	}

	config := toConfig(req.Image, req.Command, req.Env)
	hostConfig := toHostConfig(req.Resources.Memory, req.Resources.CPU, int(req.Resources.Instances), req.Resources.Volumes)

	body, err := cli.ContainerCreate(ctx, &config, &hostConfig, &network.NetworkingConfig{}, req.Name)
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	res.Status = "Container created successfully"

	return nil
}

func (*Task) Undeploy(ctx context.Context, req *proto.UndeployRequest, res *proto.UndeployResponse) error {
	res.Status = "Job" + req.Name + "undeployed"
	return nil
}

func (*Task) List(ctx context.Context, req *proto.ListRequest, res *proto.ListResponse) error {
	return nil
}

func main() {
	service := getService()

	service.Init()

	proto.RegisterTasksHandler(service.Server(), new(Task))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func getService() micro.Service {
	return micro.NewService(
		micro.Name("receptacle-agent"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
}

func getEnv() (*client.Client, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return cli, err
}

func toConfig(image string, command, env []string) container.Config {
	config := container.Config{}
	if image != "" {
		config.Image = image
	} else {
		return container.Config{}
	}
	config.Cmd = command
	config.Env = env
	return config
}

func toHostConfig(memory, CPU string, instances int, volumes []string) container.HostConfig {
	config := container.HostConfig{}

	if memory != "" {
		mem, err := unit.FromHumanSize(memory)
		config.Resources.Memory = mem
		if err != nil {
			return container.HostConfig{}
		}
	}
	if volumes != nil {
		config.Binds = volumes
	}
	if CPU != "" {
		config.CpusetMems = CPU
	}

	return config
}

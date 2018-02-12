package main

import (
	"context"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/network"
)

func CreateContainer(task Task) error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	config, err := toConfig(task.Image, task.Command, task.Env)
	if err != nil {
		return err
	}

	hostConfig, err := toHostConfig(task.Resources)
	if err != nil {
		return err
	}

	body, err := cli.ContainerCreate(ctx, &config, &hostConfig, &network.NetworkingConfig{}, task.Name)
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}

func ContainerExists(name string) (bool, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return false, err
	}

	ctx := context.Background()

	conts, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	for _, cont := range conts {
		if cont.Names[0][1:] == name {
			return true, nil
		}
	}
	return false, err
}

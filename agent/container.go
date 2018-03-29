package main

import (
	"context"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/network"
	"github.com/kristoy0/receptacle/store"
)

// CreateContainer -
// This function is used to create docker containers
// from the Task struct
func CreateContainer(task store.Task) error {
	// Fetch the docker client from environment
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Using helper functions to get docker container creation configuration
	config, err := toConfig(task.Image, task.Command, task.Env)
	if err != nil {
		return err
	}

	hostConfig, err := toHostConfig(task.Resources)
	if err != nil {
		return err
	}

	// Creating the container with given specifications
	body, err := cli.ContainerCreate(ctx, &config, &hostConfig, &network.NetworkingConfig{}, task.Name)
	if err != nil {
		return err
	}

	// Start the container
	err = cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ContainerExists -
// This function is used to check
// if a container with the given name
// already exists
func ContainerExists(name string) (bool, error) {
	// Fetch the docker client from environment
	cli, err := docker.NewEnvClient()
	if err != nil {
		return false, err
	}

	ctx := context.Background()

	// List all containers and
	// loop over them to check for a matching name
	conts, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	for _, cont := range conts {
		if cont.Names[0][1:] == name {
			return true, nil
		}
	}
	return false, err
}

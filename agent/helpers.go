package main

import (
	"errors"
	"docker.io/go-docker/api/types/container"
	unit "github.com/docker/go-units"
	"strings"
)

func toConfig(image string, command, env []string) (container.Config, error) {
	config := container.Config{}
	if image != "" {
		config.Image = image
	} else {
		return container.Config{}, errors.New("Image not specified")
	}
	config.Cmd = command
	config.Env = env
	return config, nil
}

func toHostConfig(resources Resources) (container.HostConfig, error) {
	config := container.HostConfig{}

	if resources.Memory != "" {
		mem, err := unit.FromHumanSize(resources.Memory)
		config.Resources.Memory = mem
		if err != nil {
			return container.HostConfig{}, err
		}
	}
	if resources.Volumes != nil {
		config.Binds = resources.Volumes
	}
	if resources.CPU != "" {
		config.CpusetCpus = resources.CPU
	}

	return config, nil
}

func stripDirectory(key string) (string, error){
	if strings.Contains(key, "/") {
		split := strings.Split(key, "/")
		return split[1], nil
	}
	return "", errors.New("Could not split key")
}

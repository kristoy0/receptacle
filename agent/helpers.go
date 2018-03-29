package main

import (
	"errors"
	"strings"

	"docker.io/go-docker/api/types/container"
	unit "github.com/docker/go-units"
	"github.com/kristoy0/receptacle/store"
)

// this function is used to put Task data into
// container.Config format
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

// this function is used to put Task data into
// container.HostConfig format
func toHostConfig(resources store.Resources) (container.HostConfig, error) {
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
	if resources.CPU != 0 {
		config.NanoCPUs = int64(resources.CPU * 1000000000)
	}

	config.PublishAllPorts = true

	return config, nil
}

// this function is used to
// strip the directory from a
// database key
func stripDirectory(key string) (string, error) {
	if strings.Contains(key, "/") {
		split := strings.Split(key, "/")
		return split[1], nil
	}
	return "", errors.New("Could not split key")
}

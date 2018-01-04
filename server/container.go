package server

import (
	"context"
	"errors"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	unit "github.com/docker/go-units"
	"github.com/gin-gonic/gin"
)

// Create - create a new container from task spec
func Create(c *gin.Context) {
	log.Println("Creating container")

	task := Task{}
	c.BindJSON(&task)

	cli, ctx, err := fetchEnv()
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	config, err := toConfig(task.Image, task.Command, task.Env)
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	hostConfig, err := toHostConfig(task.Resources)
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	body, err := cli.ContainerCreate(ctx, &config, &hostConfig, &network.NetworkingConfig{}, task.Name)
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	err = cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	log.Println("Container created")
	c.JSON(200, "Container successfully created")
}

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

// List - lists all containers
func List(c *gin.Context) {
	cli, ctx, err := fetchEnv()
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	res, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}
	c.JSON(200, res)
}

// Inspect - inspects a container
func Inspect(c *gin.Context) {
	cli, ctx, err := fetchEnv()
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	id := c.Param("id")
	res, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		log.Println(err)
		c.JSON(200, err)
	}

	c.JSON(200, res)
}

func fetchEnv() (*client.Client, context.Context, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, nil, err
	}

	ctx := context.Background()
	if err != nil {
		return nil, nil, err
	}
	return cli, ctx, nil
}

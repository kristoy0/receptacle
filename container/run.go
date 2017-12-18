package container

import (
	"context"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func (c Container) Run(ctx context.Context, cli *client.Client) (string, error) {
	var Ports []string
	for i := range c.Ports {
		Ports = append(Ports, c.Ports[i].IP+":"+strconv.Itoa(int(c.Ports[i].PrivatePort))+":"+strconv.Itoa(int(c.Ports[i].PublicPort)))
	}
	_, portMap, err := nat.ParsePortSpecs(Ports)
	if err != nil {
		return "", err
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:  c.ImageName,
		Env:    c.Env,
		Labels: c.Labels,
	}, &container.HostConfig{
		PortBindings:  portMap,
		RestartPolicy: container.RestartPolicy{Name: c.RestartPolicy},
		AutoRemove:    c.AutoRemove,
		Binds:         c.Volumes,
		Links:         c.Links,
	}, nil, "")
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, err
}

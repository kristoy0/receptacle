package main

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Image struct {
	Name string `json:"name,omitempty"`
}

type Container struct {
	Name          string            `json:"name,omitempty"`
	ImageName     string            `json:"image,omitempty"`
	Ports         []string          `json:"ports,omitempty"`
	Memory        string            `json:"memory,omitempty"`
	RestartPolicy string            `json:"restartPolicy,omitempty"`
	Env           []string          `json:"env,omitempty"`
	AutoRemove    bool              `json:"autoRemove,omitempty"`
	Volumes       []string          `json:"volumes,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Links         []string          `json:"links,omitempty"`
}

func main() {
	// ctx := context.Background()
	// cli, err := client.NewEnvClient()
	// if err != nil {
	// 	panic(err)
	// }

	// newImg := Image{
	// 	Name: "nginx",
	// }

	// newCont := Container{
	// 	Name:      "testcont",
	// 	ImageName: "nginx",
	// 	Ports: []string{
	// 		"0.0.0.0:8080:80",
	// 	},
	// }

	// out, err := newImg.Pull(ctx, cli)
	// if err != nil {
	// 	panic(err)
	// }

	// io.Copy(os.Stdout, out)

	// err = newCont.Run(ctx, cli)
	// if err != nil {
	// 	panic(err)
	// }
}

func (i Image) Pull(ctx context.Context, cli *client.Client) (io.ReadCloser, error) {
	out, err := cli.ImagePull(ctx, i.Name, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}

	return out, err
}

func (c Container) Run(ctx context.Context, cli *client.Client) (string, error) {
	_, portMap, err := nat.ParsePortSpecs(c.Ports)
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

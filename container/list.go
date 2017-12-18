package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func List(ctx context.Context, cli *client.Client) ([]Container, error) {
	sum, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	out := []Container{}
	for _, cont := range sum {
		out = append(out, Container{
			ID:        cont.ID,
			Name:      cont.Names[0],
			ImageName: cont.Image,
			Status:    cont.Status,
			State:     cont.State,
			Ports:     cont.Ports,
		})
	}
	return out, err
}

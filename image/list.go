package image

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	unit "github.com/docker/go-units"
)

func List(ctx context.Context, cli *client.Client) ([]Image, error) {
	sum, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		return nil, err
	}
	out := []Image{}
	for _, img := range sum {
		out = append(out, Image{
			Name: img.RepoTags[0],
			Id:   img.ID,
			Size: unit.HumanSize(float64(img.Size)),
		})
	}
	return out, err
}

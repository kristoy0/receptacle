package image

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (i Image) Pull(ctx context.Context, cli *client.Client) (io.ReadCloser, error) {
	out, err := cli.ImagePull(ctx, i.Name, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}

	return out, err
}

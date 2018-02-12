package main

import (
	"context"
	"io"
	"os"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

func PullImage(name string) error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	data, err := cli.ImagePull(ctx, name, types.ImagePullOptions{})
	// defer data.Close()
	io.Copy(os.Stdout, data)

	return nil
}

func ImageExists(name string) (bool, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return false, err
	}

	ctx := context.Background()

	conts, err := cli.ImageList(ctx, types.ImageListOptions{})
	for _, cont := range conts {
		if cont.RepoTags[0] == name {
			return true, nil
		}
	}
	return false, err
}

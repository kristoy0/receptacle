package main

import (
	"context"
	"io"
	"os"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

// PullImage -
// This function is used to pull a docker image
// given the name
func PullImage(name string) error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// this will pull the image
	data, err := cli.ImagePull(ctx, name, types.ImagePullOptions{})
	// defer data.Close()
	// output the pull status to the console
	io.Copy(os.Stdout, data)

	return nil
}

// ImageExists -
// This function is used to check if
// an image with the corresponding name
// already exists on the host
// if so the image will not be pulled again
func ImageExists(name string) (bool, error) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return false, err
	}

	ctx := context.Background()

	// Iterate over all the images on the host
	// in the look for a matching image
	conts, err := cli.ImageList(ctx, types.ImageListOptions{})
	for _, cont := range conts {
		if cont.RepoTags[0] == name {
			return true, nil
		}
	}
	return false, err
}

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/client"
	"github.com/kristoy0/receptacle-worker/container"
	"github.com/kristoy0/receptacle-worker/image"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	newImg := image.Image{
		Name: "nginx",
	}

	newCont := container.Container{
		Name:      "testcont",
		ImageName: "nginx",
		Ports: []string{
			"0.0.0.0:8080:80",
		},
	}

	out, err := newImg.Pull(ctx, cli)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	id, err := newCont.Run(ctx, cli)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

package api

import (
	"context"

	"github.com/docker/docker/client"
)

// FetchContextAndClient - fetches background context & pointer to a docker client
func FetchContextAndClient() (context.Context, *client.Client, error) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	return ctx, cli, err
}

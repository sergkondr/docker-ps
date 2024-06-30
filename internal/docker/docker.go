package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type GetOpts struct {
	SocketAddr string

	ShowAll bool
}

func GetList(ctx context.Context, opts GetOpts) ([]types.Container, error) {
	apiClient, err := client.NewClientWithOpts(
		client.WithAPIVersionNegotiation(),
		client.WithHost(opts.SocketAddr),
		client.FromEnv,
	)
	if err != nil {
		return nil, fmt.Errorf("can't initialize docker client: %v", err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(ctx, container.ListOptions{All: opts.ShowAll})
	if err != nil {
		return nil, fmt.Errorf("can't get containers list: %w", err)
	}

	return containers, nil
}

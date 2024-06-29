package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type GetOpts struct {
	SocketAddr string

	ShowAll bool
}

func GetList(ctx context.Context, opts GetOpts) ([]types.Container, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv) //client.WithHost(opts.SocketAddr))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(ctx, container.ListOptions{All: opts.ShowAll})
	if err != nil {
		return nil, fmt.Errorf("can't get containers: %w", err)
	}

	if len(containers) == 0 {
		fmt.Println("No containers running")
	}

	return containers, nil
}

func (c *ContainerInfo) String() string {
	sb := strings.Builder{}
	return sb.String()
}

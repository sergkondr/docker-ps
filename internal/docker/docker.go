package docker

import (
	"context"
	"fmt"
	"github.com/docker/cli/cli/command"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

type GetOpts struct {
	Client command.Cli

	ShowAll bool
}

func GetList(opts GetOpts) ([]types.Container, error) {
	ctx := context.Background()

	containers, err := opts.Client.Client().ContainerList(ctx, container.ListOptions{All: opts.ShowAll})
	if err != nil {
		return nil, fmt.Errorf("can't get containers list: %w", err)
	}

	return containers, nil
}

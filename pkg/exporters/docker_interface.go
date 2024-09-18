package exporters

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerInterface struct {
	dockerApiClient *client.Client
	timeout         time.Duration
}

func NewDockerInterface(ctx context.Context, timeout time.Duration) (*DockerInterface, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker API client: %w", err)
	}

	return &DockerInterface{
		dockerApiClient: apiClient,
		timeout:         timeout,
	}, nil
}

func (d *DockerInterface) collectContainerStats(ctx context.Context, containerID string) (*Statistics, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	ctx, cancel := context.WithTimeout(ctx, d.timeout)
	defer cancel()

	var statistics *Statistics
	containerStats, err := d.dockerApiClient.ContainerStats(ctx, containerID, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get container stats: %w", err)
	}
	defer func() { _ = containerStats.Body.Close() }()
	body, err := io.ReadAll(containerStats.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read container stats body: %w", err)
	}
	err = json.Unmarshal(body, &statistics)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal container stats: %w", err)
	}
	return statistics, nil
}

func (d *DockerInterface) getContainerList(ctx context.Context) ([]types.Container, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	ctx, cancel := context.WithTimeout(ctx, d.timeout)
	defer cancel()

	containers, err := d.dockerApiClient.ContainerList(ctx, container.ListOptions{All: false})
	if err != nil {
		return nil, fmt.Errorf("failed to get container list: %w", err)
	}
	return containers, nil
}

func (d *DockerInterface) getHost() string {
	return d.dockerApiClient.DaemonHost()
}

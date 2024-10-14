package exporters

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/GrzegorzMika/docker_stats_exporter/exporters/metrics"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/sync/errgroup"
)

type DockerInterface struct {
	dockerApiClient *client.Client
	timeout         time.Duration
	maxGoroutines   int
	ctx             context.Context
}

func NewDockerInterface(ctx context.Context, timeout time.Duration, maxGoroutines int) (*DockerInterface, error) {
	apiClient, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker API client: %w", err)
	}
	log.Println("Docker API client created successfully")
	return &DockerInterface{
		dockerApiClient: apiClient,
		timeout:         timeout,
		maxGoroutines:   maxGoroutines,
		ctx:             ctx,
	}, nil
}

func (d *DockerInterface) Close() error {
	return d.dockerApiClient.Close()
}

func (d *DockerInterface) CollectStats() (map[*types.Container]*metrics.Statistics, error) {
	containerList, err := d.getContainerList(d.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container list: %w", err)
	}
	statsMap := make(map[*types.Container]*metrics.Statistics, len(containerList))

	g, ctx := errgroup.WithContext(d.ctx)
	g.SetLimit(d.maxGoroutines)

	for _, container := range containerList {
		c := container
		g.Go(func() error {
			stats, err := d.collectContainerStats(ctx, container.ID)
			if err != nil {
				return fmt.Errorf("failed to collect stats for container %s: %w", c.ID, err)
			}
			statsMap[&c] = stats
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return statsMap, nil
}

func (d *DockerInterface) collectContainerStats(ctx context.Context, containerID string) (*metrics.Statistics, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	ctx, cancel := context.WithTimeout(ctx, d.timeout)
	defer cancel()

	var statistics *metrics.Statistics
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

func (d *DockerInterface) getHostName(ctx context.Context) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}
	ctx, cancel := context.WithTimeout(ctx, d.timeout)
	defer cancel()

	info, err := d.dockerApiClient.Info(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get Docker info: %w", err)
	}
	return info.Name, nil
}

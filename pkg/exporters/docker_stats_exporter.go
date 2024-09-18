package exporters

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
)

type DockerStatsExporter struct{}

func NewDockerStatsExporter(ctx context.Context) (*DockerStatsExporter, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		// fmt.Printf("%s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.Status)
		_, _ = apiClient.ContainerStats(ctx, ctr.ID, false)
	}
	return &DockerStatsExporter{}, nil
}

func (d *DockerStatsExporter) Collect(ch chan<- prometheus.Metric) {}

func (d *DockerStatsExporter) Describe(ch chan<- *prometheus.Desc) {}

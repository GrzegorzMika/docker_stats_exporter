package exporters

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type DockerStatsCollector struct {
	DockerInterface *DockerInterface
}

func NewDockerStatsCollector(ctx context.Context, reg prometheus.Registerer) (*DockerInterface, error) {
	dockerInterface, err := NewDockerInterface(ctx, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker interface: %w", err)
	}

	collector := &DockerStatsCollector{
		DockerInterface: dockerInterface,
	}

	prometheus.WrapRegistererWith(prometheus.Labels{"host": dockerInterface.getHost()}, reg).MustRegister(collector)
	return dockerInterface, nil
}

func (d *DockerStatsCollector) Collect(ch chan<- prometheus.Metric) {}

func (d *DockerStatsCollector) Describe(ch chan<- *prometheus.Desc) {}

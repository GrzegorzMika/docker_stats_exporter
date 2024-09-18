package exporters

import "github.com/prometheus/client_golang/prometheus"

type DockerStatsExporter struct{}

func NewDockerStatsExporter() *DockerStatsExporter {
	return &DockerStatsExporter{}
}

func (d *DockerStatsExporter) Collect(ch chan<- prometheus.Metric) {}

func (d *DockerStatsExporter) Describe(ch chan<- *prometheus.Desc) {}

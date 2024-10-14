package exporters

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GrzegorzMika/docker_stats_exporter/exporters/metrics"
	"github.com/docker/docker/api/types"
	"github.com/prometheus/client_golang/prometheus"
)

type metricProvider func(*types.Container, *metrics.Statistics) prometheus.Metric

type DockerStatsCollector struct {
	DockerInterface   *DockerInterface
	host              string
	metricsProviders  []metricProvider
	dockerStatsUpDesc *prometheus.Desc
}

type DockerStatsCollectorArgs struct {
	Timeout       time.Duration
	MaxGoroutines int
}

func NewDockerStatsCollector(ctx context.Context, reg prometheus.Registerer, args DockerStatsCollectorArgs) (*DockerInterface, error) {
	dockerInterface, err := NewDockerInterface(ctx, args.Timeout, args.MaxGoroutines)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker interface: %w", err)
	}

	host, err := dockerInterface.getHostName(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get host name %w", err)
	}

	metricsProviders := []metricProvider{
		readTimeProvider,
		cpuUsageTotalProvider,
		cpuSystemUsageTotalProvider,
		cpuUsageDeltaProvider,
		cpuSystemUsageDeltaProvider,
		cpuNumberProvider,
		memoryUsageProvider,
		memoryCachedUsageProvider,
		memoryLimitProvider,
		networkBytesReceivedProvider,
		networkBytesSentProvider,
		networkPacketsReceivedProvider,
		networkPacketsSentProvider,
		networkErrorsReceivedProvider,
		networkErrorsSentProvider,
	}

	dockerStatsUpDesc := prometheus.NewDesc(
		prometheus.BuildFQName("docker", "stats", "up"),
		"Whether scraping Docker Stats metrics was successful.",
		nil,
		nil,
	)
	collector := &DockerStatsCollector{
		DockerInterface:   dockerInterface,
		host:              host,
		metricsProviders:  metricsProviders,
		dockerStatsUpDesc: dockerStatsUpDesc,
	}
	log.Printf("Registering Docker metrics for host %s...\n", host)
	prometheus.WrapRegistererWith(prometheus.Labels{"host": host}, reg).MustRegister(collector)
	return dockerInterface, nil
}

func (d *DockerStatsCollector) Collect(ch chan<- prometheus.Metric) {
	containerStatistics, err := d.DockerInterface.CollectStats()
	if err != nil {
		log.Printf("Failed to collect Docker statistics: %v\n", err)
		ch <- prometheus.MustNewConstMetric(
			d.dockerStatsUpDesc,
			prometheus.GaugeValue,
			0,
		)
		return
	}
	ch <- prometheus.MustNewConstMetric(
		d.dockerStatsUpDesc,
		prometheus.GaugeValue,
		1,
	)
	for container, stats := range containerStatistics {
		for _, provider := range d.metricsProviders {
			ch <- provider(container, stats)
		}
	}
}
func (d *DockerStatsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(d, ch)
}

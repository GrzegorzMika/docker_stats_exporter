package exporters

import (
	"github.com/GrzegorzMika/docker_stats_exporter/exporters/metrics"
	"github.com/docker/docker/api/types"
	"github.com/prometheus/client_golang/prometheus"
)

func readTimeProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_read_statistics_time_seconds",
			"Last time read operation took place on a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.ReadTimeMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuUsageTotalProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_usage_seconds_total",
			"Total CPU usage for a container in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.CPUUsageTotalMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuSystemUsageTotalProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_system_usage_seconds_total",
			"Total system CPU usage in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.CPUSystemUsageTotalMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuUsageDeltaProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_usage_delta_seconds",
			"Delta CPU usage for a container in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.CPUUsageDeltaMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuSystemUsageDeltaProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_system_usage_delta_seconds",
			"Delta system CPU usage in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.CPUSystemUsageDeltaMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuNumberProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_number",
			"Number of CPUs for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.CPUNumberMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func memoryUsageProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_memory_usage_bytes_total",
			"Memory usage for a container in bytes including cache",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.MemoryUsageMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func memoryCachedUsageProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_memory_cached_usage_bytes_total",
			"Memory usage for a container as cache",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.MemoryCachedUsageMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func memoryLimitProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_memory_limit_bytes_total",
			"Memory limit for a container in bytes",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.MemoryLimitMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func networkBytesReceivedProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_network_bytes_received_bytes_total",
			"Network bytes received for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.NetworkBytesReceivedMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func networkBytesSentProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_network_bytes_sent_bytes_total",
			"Network bytes sent for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.NetworkBytesSentMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func networkPacketsReceivedProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_network_packets_received_total",
			"Network packets received for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.NetworkPacketsReceivedMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func networkPacketsSentProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_network_packets_sent_total",
			"Network packets sent for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.NetworkPacketsSentMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func networkErrorsReceivedProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_network_errors_received_total",
			"Network errors received for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.NetworkErrorsReceivedMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func networkErrorsSentProvider(container *types.Container, stats *metrics.Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_network_errors_sent_total",
			"Network errors sent for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		metrics.NetworkErrorsSentMetric(stats),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

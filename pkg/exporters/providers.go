package exporters

import (
	"github.com/docker/docker/api/types"
	"github.com/prometheus/client_golang/prometheus"
)

func readTimeProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_read_statistics_time_seconds",
			"Last time read operation took place on a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.Read.UnixMicro())/1e6,
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuUsageTotalProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_usage_seconds_total",
			"Total CPU usage for a container in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.CPUStats.CPUUsage.TotalUsage),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuSystemUsageTotalProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_system_usage_seconds_total",
			"Total system CPU usage in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.CPUStats.SystemCPUUsage),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuUsageDeltaProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_usage_delta_seconds",
			"Delta CPU usage for a container in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.CPUStats.CPUUsage.TotalUsage)-float64(stats.PrecpuStats.CPUUsage.TotalUsage),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuSystemUsageDeltaProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_system_usage_delta_seconds",
			"Delta system CPU usage in seconds",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.CPUStats.SystemCPUUsage)-float64(stats.PrecpuStats.SystemCPUUsage),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func cpuNumberProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_cpu_number",
			"Number of CPUs for a container",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(len(stats.CPUStats.CPUUsage.PercpuUsage)),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func memoryUsageProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_memory_usage_bytes",
			"Memory usage for a container in bytes including cache",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.MemoryStats.Usage),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func memoryCachedUsageProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_memory_cached_usage_bytes",
			"Memory usage for a container as cache",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.MemoryStats.Stats.Cache),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

func memoryLimitProvider(container *types.Container, stats *Statistics) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"docker_container_memory_limit_bytes",
			"Memory limit for a container in bytes",
			[]string{"container_id", "container_name", "image_id", "image_name"},
			nil,
		),
		prometheus.GaugeValue,
		float64(stats.MemoryStats.Limit),
		container.ID, container.Names[0], container.ImageID, container.Image,
	)
}

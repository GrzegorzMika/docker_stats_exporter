package metrics

import "slices"

func ReadTimeMetric(stats *Statistics) float64 {
	return float64(stats.Read.UnixMicro()) / 1e6
}

func CPUUsageTotalMetric(stats *Statistics) float64 {
	return float64(stats.CPUStats.CPUUsage.TotalUsage)
}

func CPUSystemUsageTotalMetric(stats *Statistics) float64 {
	return float64(stats.CPUStats.SystemCPUUsage)
}

func CPUUsageDeltaMetric(stats *Statistics) float64 {
	return float64(stats.CPUStats.CPUUsage.TotalUsage) - float64(stats.PrecpuStats.CPUUsage.TotalUsage)
}

func CPUSystemUsageDeltaMetric(stats *Statistics) float64 {
	return float64(stats.CPUStats.SystemCPUUsage) - float64(stats.PrecpuStats.SystemCPUUsage)
}

func CPUNumberMetric(stats *Statistics) float64 {
	return slices.Max([]float64{float64(stats.CPUStats.OnlineCpus), float64(stats.PrecpuStats.OnlineCpus), float64(len(stats.CPUStats.CPUUsage.PercpuUsage))})
}

func MemoryUsageMetric(stats *Statistics) float64 {
	return float64(stats.MemoryStats.Usage)
}

func MemoryCachedUsageMetric(stats *Statistics) float64 {
	return float64(stats.MemoryStats.Stats.Cache)
}

func MemoryLimitMetric(stats *Statistics) float64 {
	return float64(stats.MemoryStats.Limit)
}

func NetworkBytesReceivedMetric(stats *Statistics) float64 {
	return float64(stats.Networks.Eth0.RxBytes)
}

func NetworkBytesSentMetric(stats *Statistics) float64 {
	return float64(stats.Networks.Eth0.TxBytes)
}

func NetworkPacketsReceivedMetric(stats *Statistics) float64 {
	return float64(stats.Networks.Eth0.RxPackets)
}

func NetworkPacketsSentMetric(stats *Statistics) float64 {
	return float64(stats.Networks.Eth0.TxPackets)
}

func NetworkErrorsReceivedMetric(stats *Statistics) float64 {
	return float64(stats.Networks.Eth0.RxErrors)
}

func NetworkErrorsSentMetric(stats *Statistics) float64 {
	return float64(stats.Networks.Eth0.TxErrors)
}

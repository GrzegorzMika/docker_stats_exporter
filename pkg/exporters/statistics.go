package exporters

import "time"

type Statistics struct {
	Read        time.Time   `json:"read"`
	PidsStats   PidsStats   `json:"pids_stats"`
	Networks    Networks    `json:"networks"`
	MemoryStats MemoryStats `json:"memory_stats"`
	BlkioStats  BlkioStats  `json:"blkio_stats"`
	CPUStats    CPUStats    `json:"cpu_stats"`
	PrecpuStats PrecpuStats `json:"precpu_stats"`
}
type PidsStats struct {
	Current int `json:"current"`
}
type Eth0 struct {
	RxBytes   int `json:"rx_bytes"`
	RxDropped int `json:"rx_dropped"`
	RxErrors  int `json:"rx_errors"`
	RxPackets int `json:"rx_packets"`
	TxBytes   int `json:"tx_bytes"`
	TxDropped int `json:"tx_dropped"`
	TxErrors  int `json:"tx_errors"`
	TxPackets int `json:"tx_packets"`
}
type Eth5 struct {
	RxBytes   int `json:"rx_bytes"`
	RxDropped int `json:"rx_dropped"`
	RxErrors  int `json:"rx_errors"`
	RxPackets int `json:"rx_packets"`
	TxBytes   int `json:"tx_bytes"`
	TxDropped int `json:"tx_dropped"`
	TxErrors  int `json:"tx_errors"`
	TxPackets int `json:"tx_packets"`
}
type Networks struct {
	Eth0 Eth0 `json:"eth0"`
	Eth5 Eth5 `json:"eth5"`
}
type Stats struct {
	TotalPgmajfault         int `json:"total_pgmajfault"`
	Cache                   int `json:"cache"`
	MappedFile              int `json:"mapped_file"`
	TotalInactiveFile       int `json:"total_inactive_file"`
	Pgpgout                 int `json:"pgpgout"`
	Rss                     int `json:"rss"`
	TotalMappedFile         int `json:"total_mapped_file"`
	Writeback               int `json:"writeback"`
	Unevictable             int `json:"unevictable"`
	Pgpgin                  int `json:"pgpgin"`
	TotalUnevictable        int `json:"total_unevictable"`
	Pgmajfault              int `json:"pgmajfault"`
	TotalRss                int `json:"total_rss"`
	TotalRssHuge            int `json:"total_rss_huge"`
	TotalWriteback          int `json:"total_writeback"`
	TotalInactiveAnon       int `json:"total_inactive_anon"`
	RssHuge                 int `json:"rss_huge"`
	HierarchicalMemoryLimit int `json:"hierarchical_memory_limit"`
	TotalPgfault            int `json:"total_pgfault"`
	TotalActiveFile         int `json:"total_active_file"`
	ActiveAnon              int `json:"active_anon"`
	TotalActiveAnon         int `json:"total_active_anon"`
	TotalPgpgout            int `json:"total_pgpgout"`
	TotalCache              int `json:"total_cache"`
	InactiveAnon            int `json:"inactive_anon"`
	ActiveFile              int `json:"active_file"`
	Pgfault                 int `json:"pgfault"`
	InactiveFile            int `json:"inactive_file"`
	TotalPgpgin             int `json:"total_pgpgin"`
}
type MemoryStats struct {
	Stats    Stats `json:"stats"`
	MaxUsage int   `json:"max_usage"`
	Usage    int   `json:"usage"`
	Failcnt  int   `json:"failcnt"`
	Limit    int   `json:"limit"`
}
type BlkioStats struct {
}
type CPUUsage struct {
	PercpuUsage       []int `json:"percpu_usage"`
	UsageInUsermode   int   `json:"usage_in_usermode"`
	TotalUsage        int   `json:"total_usage"`
	UsageInKernelmode int   `json:"usage_in_kernelmode"`
}
type ThrottlingData struct {
	Periods          int `json:"periods"`
	ThrottledPeriods int `json:"throttled_periods"`
	ThrottledTime    int `json:"throttled_time"`
}
type CPUStats struct {
	CPUUsage       CPUUsage       `json:"cpu_usage"`
	SystemCPUUsage int64          `json:"system_cpu_usage"`
	OnlineCpus     int            `json:"online_cpus"`
	ThrottlingData ThrottlingData `json:"throttling_data"`
}
type PrecpuStats struct {
	CPUUsage       CPUUsage       `json:"cpu_usage"`
	SystemCPUUsage int64          `json:"system_cpu_usage"`
	OnlineCpus     int            `json:"online_cpus"`
	ThrottlingData ThrottlingData `json:"throttling_data"`
}

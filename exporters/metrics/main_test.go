package metrics

import (
	"encoding/json"
	"os"
	"testing"
)

var example_response *Statistics

var example_response_raw = `{
  "read": "2024-09-18T15:03:20.129723697Z",
  "preread": "2024-09-18T15:03:19.126780933Z",
  "pids_stats": {
    "current": 159,
    "limit": 37705
  },
  "blkio_stats": {
    "io_service_bytes_recursive": [
      {
        "major": 259,
        "minor": 0,
        "op": "read",
        "value": 1103695872
      },
      {
        "major": 259,
        "minor": 0,
        "op": "write",
        "value": 16975872
      }
    ],
    "io_serviced_recursive": null,
    "io_queue_recursive": null,
    "io_service_time_recursive": null,
    "io_wait_time_recursive": null,
    "io_merged_recursive": null,
    "io_time_recursive": null,
    "sectors_recursive": null
  },
  "num_procs": 0,
  "storage_stats": {},
  "cpu_stats": {
    "cpu_usage": {
      "total_usage": 20281474000,
      "usage_in_kernelmode": 7590287000,
      "usage_in_usermode": 12691186000
    },
    "system_cpu_usage": 13020180000000,
    "online_cpus": 16,
    "throttling_data": {
      "periods": 0,
      "throttled_periods": 0,
      "throttled_time": 0
    }
  },
  "precpu_stats": {
    "cpu_usage": {
      "total_usage": 20260627000,
      "usage_in_kernelmode": 7582485000,
      "usage_in_usermode": 12678141000
    },
    "system_cpu_usage": 13004210000000,
    "online_cpus": 16,
    "throttling_data": {
      "periods": 0,
      "throttled_periods": 0,
      "throttled_time": 0
    }
  },
  "memory_stats": {
    "usage": 1975832576,
    "stats": {
      "active_anon": 874655744,
      "active_file": 25579520,
      "anon": 874713088,
      "anon_thp": 0,
      "file": 1087025152,
      "file_dirty": 0,
      "file_mapped": 21024768,
      "file_writeback": 0,
      "inactive_anon": 0,
      "inactive_file": 1061445632,
      "kernel_stack": 2605056,
      "pgactivate": 0,
      "pgdeactivate": 0,
      "pgfault": 455452,
      "pglazyfree": 0,
      "pglazyfreed": 0,
      "pgmajfault": 219,
      "pgrefill": 0,
      "pgscan": 0,
      "pgsteal": 0,
      "shmem": 0,
      "slab": 6655968,
      "slab_reclaimable": 3763152,
      "slab_unreclaimable": 2892816,
      "sock": 0,
      "thp_collapse_alloc": 0,
      "thp_fault_alloc": 0,
      "unevictable": 0,
      "workingset_activate": 0,
      "workingset_nodereclaim": 0,
      "workingset_refault": 0
    },
    "limit": 33013448704
  },
  "name": "/mssql",
  "id": "04d4f7b337737bd8b750783e44ebb5e37c2e9a57de48572432349df811ff5002",
  "networks": {
    "eth0": {
      "rx_bytes": 88452,
      "rx_packets": 1944,
      "rx_errors": 0,
      "rx_dropped": 0,
      "tx_bytes": 1786,
      "tx_packets": 25,
      "tx_errors": 0,
      "tx_dropped": 0
    }
  }
}
`

func parseResponse(response string) (*Statistics, error) {
	var stats Statistics
	err := json.Unmarshal([]byte(response), &stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

func TestMain(m *testing.M) {
	var err error
	example_response, err = parseResponse(example_response_raw)
	if err != nil {
		panic("failed to parse example response: " + err.Error())
	}
	os.Exit(m.Run())
}

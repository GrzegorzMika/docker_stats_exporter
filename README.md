# Docker Stats exporter for Prometheus

Prometheus exporter for Docker containers resource usage statistics. It converts the output from Docker API [`/containers/(id)/stats`](https://docs.docker.com/reference/api/engine/version/v1.46/#tag/Container/operation/ContainerStats) endpoint into a Prometheus-compatible format. The exporter automatically detects all Docker containers and provides the metrics for them.

## Building and running the exporter
### Build and run locally
```bash
git clone git@github.com:GrzegorzMika/docker_stats_exporter.git
cd docker_stats_exporter
go build .
./docker_stats_exporter -version
```

### Pre-build binaries
For pre-built binaries please take a look at [the releases](https://github.com/GrzegorzMika/docker_stats_exporter/releases).

### Running with Docker
```bash
docker run \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -p 9273:9273 \
    ghcr.io/grzegorzmika/docker_stats_exporter:latest
```

Exporter Docker image is also available via Docker Hub `gregmika/docker_stats_exporter`.

### Test with
```bash
curl http://localhost:9273/metrics
```

## Basic Prometheus Configuration
Add a block to the `scrape_configs` of your `prometheus.yml` config file:
```yaml
scrape_configs:
  - job_name: docker_stats_exporter
    static_configs:
    - targets: ['<<DOCKER-STATS-EXPORTER-HOSTNAME>>:9273']
```
and adjust the host name accordingly.

## Command-line flags
| **Name**                         | **Description** | **Default value**
|----------------------------------|-----------------|-------------------
| web.listen-address               | Address to listen on for web interface and telemetry. | `:9273`
| web.telemetry-path               | Path under which to expose metrics. | `/metrics`
| web.timeout                      | Docker Stats Exporter API request timeout. | `5s`
| internal.max-concurrent-requests | Maximum number of concurrent Docker API requests. | `10`
| version                          | Show version information and exit. |

## Available metrics
- `docker_stats_up` - Whether scraping Docker Stats metrics was successful.
- `docker_container_read_statistics_time_seconds` - Last time read operation took place on a container
- `docker_container_cpu_usage_seconds_total` - Total CPU usage for a container in seconds
- `docker_container_cpu_system_usage_seconds_total` - Total system CPU usage in seconds
- `docker_container_cpu_usage_delta_seconds` - Delta CPU usage for a container in seconds
- `docker_container_cpu_system_usage_delta_seconds` - Delta system CPU usage in seconds
- `docker_container_cpu_number` - Number of CPUs for a container
- `docker_container_memory_usage_bytes_total` - Memory usage for a container in bytes including cache
- `docker_container_memory_cached_usage_bytes_total` - Memory usage for a container as cache
- `docker_container_memory_limit_bytes_total` - Memory limit for a container in bytes
- `docker_container_network_bytes_received_bytes_total` - Network bytes received for a container
- `docker_container_network_bytes_sent_bytes_total` - Network bytes sent for a container
- `docker_container_network_packets_received_total` - Network packets received for a container
- `docker_container_network_packets_sent_total` - Network packets sent for a container
- `docker_container_network_errors_received_total` - Network errors received for a container
- `docker_container_network_errors_sent_total` - Network errors sent for a container

The measurements are labeled with container id, container name, image id, image name and host name.

## Tips and tricks

To calculate the values shown by the `docker stats` command of the docker cli tool the following formulas can be used:
- CPU usage %
```bash
(docker_container_cpu_usage_delta_seconds / docker_container_cpu_system_usage_delta_seconds) * docker_container_cpu_number * 100.0
```
- Memory usage %
```bash
(docker_container_memory_usage_bytes_total - docker_container_memory_cached_usage_bytes_total) / docker_container_memory_limit_bytes_total * 100.0
```
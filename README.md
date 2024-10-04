# Docker Stats exporter for Prometheus

Prometheus exporter for Docker containers resource usage statistics. It converts the output from Docker API [`/containers/(id)/stats`](https://docs.docker.com/reference/api/engine/version/v1.46/#tag/Container/operation/ContainerStats) endpoint into a Prometheus-compatible format. The exporter automatically detects all Docker containers and provides the metrics for them.
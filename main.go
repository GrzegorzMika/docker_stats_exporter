// Copyright 2024 Grzegorz Mika

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/GrzegorzMika/docker_stats_exporter/exporters"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultAPITimeout    = 5 * time.Second
	defaultMaxGoroutines = 10
)

var (
	version = "<<< filled in by build >>>"
	commit  = "<<< filled in by build >>>"
	date    = "<<< filled in by build >>>"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		listenAddress = flag.String("web.listen-address", ":9273", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		timeout       = flag.Duration("web.timeout", defaultAPITimeout, "Docker Stats Exporter API request timeout.")
		maxGoroutines = flag.Int("internal.max-concurrent-requests", defaultMaxGoroutines, "Maximum number of concurrent Docker API requests.")
		showVersion   = flag.Bool("version", false, "Show version information and exit.")
	)
	flag.Parse()

	log.Println(getVersion())
	if *showVersion {
		os.Exit(0)
	}

	log.Printf("Starting Docker Stats Exporter\n")
	log.Printf("API timeout: %v\n", *timeout)
	log.Printf("Max Goroutines: %v\n", *maxGoroutines)

	log.Println("Creating registry...")
	reg := prometheus.NewPedanticRegistry()
	log.Println("Creating Docker stats collector...")
	dockerAPIClient, err := exporters.NewDockerStatsCollector(
		ctx,
		reg,
		exporters.DockerStatsCollectorArgs{
			Timeout:       *timeout,
			MaxGoroutines: *maxGoroutines,
		})
	if err != nil {
		log.Panicf("Error creating Docker stats collector: %v", err)
	}

	defer func() { _ = dockerAPIClient.Close() }()

	// Add the standard process and Go metrics to the custom registry.
	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	log.Printf("Starting HTTP server on %s...\n", *listenAddress)
	http.Handle(*metricsPath, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`
			<html>
			<head><title>Docker Stats Exporter</title></head>
			<body>
			<h1>Docker Stats exporter for Prometheus</h1>
			<p><a href='` + *metricsPath + `'>Metrics</a></p>
			</body>
			</html>`))
		if err != nil {
			log.Printf("Error writing HTML: %v", err)
		}
	})
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func getVersion() string {
	return fmt.Sprintf("Version: %s, Commit SHA: %s, Build Date: %s, Go Version: %s", version, commit, date, runtime.Version())
}

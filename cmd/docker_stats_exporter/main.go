// Copyright 2024 Grzegorz Mika

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/GrzegorzMika/docker_stats_exporter/pkg/exporters"
	"github.com/GrzegorzMika/docker_stats_exporter/pkg/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const defaultAPITimeout = 5 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		listenAddress = flag.String("web.listen-address", ":9273", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		timeout       = flag.Duration("timeout", defaultAPITimeout, "API request timeout")
		showVersion   = flag.Bool("version", false, "Show version information and exit")
	)
	flag.Parse()

	log.Println(version.GetVersion())
	if *showVersion {
		os.Exit(0)
	}

	log.Printf("Starting Docker Stats Exporter\n")
	log.Printf("Listen address: %v\n", *listenAddress)
	log.Printf("API timeout: %v\n", *timeout)

	reg := prometheus.NewPedanticRegistry()
	_, err := exporters.NewDockerStatsCollector(ctx, reg)
	if err != nil {
		log.Fatalf("Error creating Docker stats collector: %v", err)
	}

	// Add the standard process and Go metrics to the custom registry.
	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	http.Handle(*metricsPath, promhttp.Handler())
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

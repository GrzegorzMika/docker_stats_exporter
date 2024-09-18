// Copyright 2024 Grzegorz Mika

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/GrzegorzMika/docker_stats_exporter/pkg/exporters"
	"github.com/GrzegorzMika/docker_stats_exporter/pkg/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":9273", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		showVersion   = flag.Bool("version", false, "Show version information and exit")
	)
	flag.Parse()

	log.Println(version.GetVersion())
	if *showVersion {
		os.Exit(0)
	}

	log.Printf("Starting Docker Stats Exporter\n")
	log.Printf("Listen address: %v\n", *listenAddress)

	exporter := exporters.NewDockerStatsExporter()

	prometheus.MustRegister(exporter)

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

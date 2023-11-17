package telemetry

// Inspired by the example here:
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gnolang/gno/telemetry/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

const meterName = "gno.land"

var enabled bool

func IsEnabled() bool {
	return enabled
}

func Init(ctx context.Context) error {
	enabled = true
	// The exporter embeds a default OpenTelemetry Reader and
	// implements prometheus.Collector, allowing it to be used as
	// both a Reader and Collector.
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}

	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(meterName)

	// Start the prometheus HTTP server and pass the exporter Collector to it
	go serveMetrics()

	// Initialize metrics to be collected.
	return metrics.Init(ctx, meter)
}

func serveMetrics() {
	log.Printf("serving metrics at localhost:4591/metrics")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":4591", nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}

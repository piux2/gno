package telemetry

// Inspired by the example here:
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gnolang/gno/telemetry/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

const (
	meterName          = "gno.land"
	defaultPort uint64 = 4591
)

var enabled bool

func IsEnabled() bool {
	return enabled
}

func Init(ctx context.Context, port uint64) error {
	enabled = true

	if port == 0 {
		port = defaultPort
	}

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
	go serveMetrics(ctx, port)

	// Initialize metrics to be collected.
	return metrics.Init(ctx, meter)
}

func serveMetrics(ctx context.Context, port uint64) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	server := http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(port), 10),
		Handler: mux,
		// Nothing should need a connection for longer than a few seconds when scraping metrics.
		BaseContext: func(net.Listener) context.Context {
			boundedCtx, _ := context.WithTimeout(ctx, time.Second*10)
			return boundedCtx
		},
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("error serving metrics over http: %v", err)
		return
	}
}

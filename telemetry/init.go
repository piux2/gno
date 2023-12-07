package telemetry

// Inspired by the example here:
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gnolang/gno/telemetry/metrics"
	"github.com/gnolang/gno/telemetry/traces"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
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
	// exporter, err := prometheus.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// provider := metric.NewMeterProvider(metric.WithReader(exporter))
	// meter := provider.Meter(meterName)

	// Start the prometheus HTTP server and pass the exporter Collector to it
	// go serveMetrics(ctx, port)

	// Use oltp metric exporter
	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

	provider := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(exporter)))
	meter := provider.Meter(meterName)

	// otel.SetMeterProvider(meterProvider)

	// Initialize metrics to be collected.
	if err := metrics.Init(ctx, meter); err != nil {
		return err
	}

	// Tracing initialization.
	_ = traces.Init()

	return nil
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

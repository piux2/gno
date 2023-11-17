package metrics

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

var (
	ctx context.Context

	// Metrics.
	BroadcastTxHistogram Int64Histogram
)

func Init(setCtx context.Context, meter metric.Meter) error {
	ctx = setCtx

	broadcastTxHistogram, err := meter.Int64Histogram(
		"broadcast_tx_hist",
		metric.WithDescription("a histogram for recording broadcast tx duration"),
		metric.WithExplicitBucketBoundaries(0, 16, 32, 64, 128, 256, 512),
		metric.WithUnit("ms"),
	)
	if err != nil {
		return err
	}
	BroadcastTxHistogram.Int64Histogram = broadcastTxHistogram

	return nil
}

type Int64Histogram struct {
	metric.Int64Histogram
}

func (h Int64Histogram) Record(value int64) {
	h.Int64Histogram.Record(ctx, value)
}

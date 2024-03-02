package telemetry

import "github.com/gnolang/gno/telemetry/options"

type Option func(*options.Config)

func WithOptionMetricsEnabled() Option {
	return func(c *options.Config) {
		c.MetricsEnabled = true
	}
}

func WithOptionTracesEnabled() Option {
	return func(c *options.Config) {
		c.TracesEnabled = true
	}
}

func WithOptionPort(port uint64) Option {
	return func(c *options.Config) {
		if port != 0 {
			c.Port = port
		}
	}
}

func WithOptionMeterName(meterName string) Option {
	return func(c *options.Config) {
		if meterName != "" {
			c.MeterName = meterName
		}
	}
}

func WithOptionExporterEndpoint(exporterEndpoint string) Option {
	return func(c *options.Config) {
		if exporterEndpoint != "" {
			c.ExporterEndpoint = exporterEndpoint
		}
	}
}

func WithOptionFakeMetrics() Option {
	return func(c *options.Config) {
		c.UseFakeMetrics = true
	}
}

func WithOptionServiceName(serviceName string) Option {
	return func(c *options.Config) {
		if serviceName != "" {
			c.ServiceName = serviceName
		}
	}
}

func WithOptionTraceFilter(traceType int64) Option {
	return func(c *options.Config) {
		c.TraceType = traceType
	}
}

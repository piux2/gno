package traces

import (
	"context"

	"github.com/gnolang/gno/tm2/pkg/sdk"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type SpanEnder struct {
	span trace.Span
}

func (s *SpanEnder) End() {
	if s != nil {
		s.span.End()
	}
}

func StartSpan(
	ctx sdk.Context,
	name string,
	attributes ...attribute.KeyValue,
) (sdk.Context, *SpanEnder) {
	spanCtx, span := otel.GetTracerProvider().Tracer("gno.land").Start(
		ctx.Context(),
		name,
		trace.WithAttributes(attributes...),
	)
	return ctx.WithContext(spanCtx), &SpanEnder{span: span}
}

// TODO: dry it
func StartSpanWithStdCtx(
	ctx context.Context,
	name string,
	attributes ...attribute.KeyValue,
) (context.Context, *SpanEnder) {
	spanCtx, span := otel.GetTracerProvider().Tracer("gno.land").Start(
		ctx,
		name,
		trace.WithAttributes(attributes...),
	)
	return spanCtx, &SpanEnder{span: span}
}

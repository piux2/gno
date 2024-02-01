package traces

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type SpanEnder struct {
	goroutineID        int
	parentNamespaceCtx namespaceContext
	span               trace.Span
}

func (s *SpanEnder) End() {
	if s == nil {
		return
	}

	namespaces[s.goroutineID] = s.parentNamespaceCtx
	s.span.End()
}

func StartSpan(
	_ namespace,
	name string,
	attributes ...attribute.KeyValue,
) *SpanEnder {
	id := goroutineID()
	parentNamespaceCtx := namespaces[id]

	spanCtx, span := otel.GetTracerProvider().Tracer("gno.land").Start(
		parentNamespaceCtx.ctx,
		name,
		trace.WithAttributes(attribute.String("component", string(parentNamespaceCtx.namespace))),
		trace.WithAttributes(attributes...),
	)

	spanEnder := &SpanEnder{
		goroutineID:        id,
		parentNamespaceCtx: parentNamespaceCtx,
		span:               span,
	}

	namespaces[id] = namespaceContext{namespace: parentNamespaceCtx.namespace, ctx: spanCtx}
	return spanEnder
}

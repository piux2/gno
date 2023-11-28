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
	// if !ok {
	// 	// Special case that happens during VM initialization. Change the namespace to vmInit.
	// 	namespace = NamespaceVMInit
	// 	parentCtx = namespaces[namespace]
	// }

	// if s := trace.SpanFromContext(parentCtx); s != nil && s.IsRecording() {
	// 	fmt.Println("still recording and overwriting")
	// }

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

// func StartSpanWithSDKCtx(
// 	ctx sdk.Context,
// 	name string,
// 	attributes ...attribute.KeyValue,
// ) (sdk.Context, *SpanEnder) {
// 	spanCtx, span := otel.GetTracerProvider().Tracer("gno.land").Start(
// 		ctx.Context(),
// 		name,
// 		trace.WithAttributes(attributes...),
// 	)

// 	return ctx.WithContext(spanCtx), &SpanEnder{span: span}
// }

// // TODO: dry it
// func StartSpanWithStdCtx(
// 	ctx context.Context,
// 	name string,
// 	attributes ...attribute.KeyValue,
// ) (context.Context, *SpanEnder) {
// 	spanCtx, span := otel.GetTracerProvider().Tracer("gno.land").Start(
// 		ctx,
// 		name,
// 		trace.WithAttributes(attributes...),
// 	)
// 	return spanCtx, &SpanEnder{span: span}
// }

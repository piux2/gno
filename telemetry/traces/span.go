package traces

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type SpanEnder struct {
	parentCtx context.Context
	namespace namespace
	span      trace.Span
}

func (s *SpanEnder) End() {
	if s == nil {
		return
	}

	namespaces[s.namespace] = s.parentCtx
	s.span.End()
}

func StartSpan(
	namespace namespace,
	name string,
	attributes ...attribute.KeyValue,
) *SpanEnder {
	parentCtx, ok := namespaces[namespace]
	if !ok {
		// Special case that happens during VM initialization. Change the namespace to vmInit.
		namespace = NamespaceVMInit
		parentCtx = namespaces[namespace]
	}

	// if s := trace.SpanFromContext(parentCtx); s != nil && s.IsRecording() {
	// 	fmt.Println("still recording and overwriting")
	// }

	spanCtx, span := otel.GetTracerProvider().Tracer("gno.land").Start(
		parentCtx,
		name,
		trace.WithAttributes(attribute.String("component", string(namespace))),
		trace.WithAttributes(attributes...),
	)

	spanEnder := &SpanEnder{
		parentCtx: parentCtx,
		namespace: namespace,
		span:      span,
	}

	namespaces[namespace] = spanCtx
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

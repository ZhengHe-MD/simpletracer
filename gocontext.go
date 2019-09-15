// codes in this file originates from https://github.com/opentracing/opentracing-go/blob/master/gocontext.go
package simpletracer

import (
	"context"
	. "intracer/api"
)

type contextKey struct{}

var activeSpanKey = contextKey{}

func ContextWithSpan(ctx context.Context, span Span) context.Context {
	return context.WithValue(ctx, activeSpanKey, span)
}

func SpanFromContext(ctx context.Context) Span {
	val := ctx.Value(activeSpanKey)
	if sp, ok := val.(Span); ok {
		return sp
	}
	return nil
}

func StartSpanFromContext(ctx context.Context, operationName string, opts ...StartSpanOption) (Span, context.Context) {
	return StartSpanFromContextWithTracer(ctx, GlobalTracer(), operationName, opts...)
}

func StartSpanFromContextWithTracer(ctx context.Context, tracer Tracer, operationName string, opts ...StartSpanOption) (Span, context.Context) {
	if parentSpan := SpanFromContext(ctx); parentSpan != nil {
		opts = append(opts, ChildOf(parentSpan.Context()))
	}
	span := tracer.StartSpan(operationName, opts...)
	return span, ContextWithSpan(ctx, span)
}
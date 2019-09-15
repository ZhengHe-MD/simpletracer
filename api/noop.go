package api

type NoopTracer struct{}

type noopSpan struct{}
type noopSpanContext struct{}

var (
	defaultNoopSpanContext SpanContext = noopSpanContext{}
	defaultNoopSpan        Span        = noopSpan{}
	defaultNoopTracer      Tracer      = NoopTracer{}
)

func (n noopSpan) Context() SpanContext                                  { return defaultNoopSpanContext }
func (n noopSpan) Finish()                                               {}
func (n noopSpan) SetOperationName(operationName string) Span            { return n }
func (n noopSpan) Tracer() Tracer                                        { return defaultNoopTracer }

// StartSpan belongs to the Tracer interface.
func (n NoopTracer) StartSpan(operationName string, opts ...StartSpanOption) Span {
	return defaultNoopSpan
}
package api

// SpanContext represents Span state that must propagate to descendant Spans
type SpanContext interface {}

// Span represents an active, un-finished span in the OpenTracing system.
//
// Spans are created by the Tracer interface.
type Span interface {
	// Sets the end timestamp and finalizes Span state.
	//
	// With the exception of calls to Context() (which are always allowed),
	// Finish() must be the last call made to any span instance, and to do
	// otherwise leads to undefined behavior.
	Finish()

	// Context() yields the SpanContext for this Span. Note that the return
	// value of Context() is still valid after a call to Span.Finish(), as is
	// a call to Span.Context() after a call to Span.Finish().
	Context() SpanContext

	// Sets or changes the operation name.
	//
	// Returns a reference to this Span for chaining.
	SetOperationName(operationName string) Span

	// Provides access to the Tracer that created this Span.
	Tracer() Tracer
}


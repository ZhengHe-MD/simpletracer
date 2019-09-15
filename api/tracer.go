package api

// Tracer is a simple, thin interface for Span creation
type Tracer interface {
	StartSpan(operationName string, opts ...StartSpanOption) Span
}

type StartSpanOptions struct {
	References []SpanReference
}

type StartSpanOption interface {
	Apply(*StartSpanOptions)
}

type SpanReferenceType int

const (
	ChildOfRef SpanReferenceType = iota
	FollowsFromRef
)

type SpanReference struct {
	Type              SpanReferenceType
	ReferencedContext SpanContext
}

// Apply satisfies the StartSpanOption interface.
func (r SpanReference) Apply(o *StartSpanOptions) {
	if r.ReferencedContext != nil {
		o.References = append(o.References, r)
	}
}

// ChildOf returns a StartSpanOption pointing to a dependent parent span.
// If sc == nil, the option has no effect.
//
// See ChildOfRef, SpanReference
func ChildOf(sc SpanContext) SpanReference {
	return SpanReference{
		Type:              ChildOfRef,
		ReferencedContext: sc,
	}
}

// FollowsFrom returns a StartSpanOption pointing to a parent Span that caused
// the child Span but does not directly depend on its result in any way.
// If sc == nil, the option has no effect.
//
// See FollowsFromRef, SpanReference
func FollowsFrom(sc SpanContext) SpanReference {
	return SpanReference{
		Type:              FollowsFromRef,
		ReferencedContext: sc,
	}
}

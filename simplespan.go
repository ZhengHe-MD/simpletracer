package simpletracer

import (
	"fmt"
	. "intracer/api"
	"strings"
	"sync"
	"time"
)

const RootParentID = uint64(0)

type SimpleSpanContext struct {
	TraceID uint64
	SpanID  uint64
}

func (i SimpleSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {
	panic("implement me")
}

type SimpleSpan struct {
	mu sync.RWMutex

	spanContext SimpleSpanContext
	parentID    uint64

	operationName string
	startTime     time.Time
	finishTime    time.Time

	tracer *SimpleTracer
}

func newInSpan(t *SimpleTracer, operationName string, opts StartSpanOptions) *SimpleSpan {
	traceID := t.randomID()
	parentID := RootParentID

	if len(opts.References) > 0 {
		parentRefContext := opts.References[0].ReferencedContext.(SimpleSpanContext)
		traceID = parentRefContext.TraceID
		parentID = parentRefContext.SpanID
	}

	spanContext := SimpleSpanContext{
		TraceID: traceID,
		SpanID:  t.randomID(),
	}

	startTime := time.Now()
	return &SimpleSpan{
		spanContext:   spanContext,
		parentID:      parentID,
		operationName: operationName,
		startTime:     startTime,
		tracer:        t,
	}
}

func (s *SimpleSpan) Finish() {
	s.finishTime = time.Now()
	if !s.isRoot() {
		return
	}

	var curLevel = []*SimpleSpan{s}
	var parts []string
	s.tracer.mu.Lock()
	for len(curLevel) > 0 {
		var nextLevel []*SimpleSpan
		for _, ss := range curLevel {
			delete(s.tracer.spanIDToSpan, ss.spanContext.SpanID)
			parts = append(parts, fmt.Sprintf("%s:%s",
				ss.operationName, ss.finishTime.Sub(ss.startTime)))
			children := s.tracer.parentIDToSpans[ss.spanContext.SpanID]
			nextLevel = append(nextLevel, children...)
			delete(s.tracer.parentIDToSpans, ss.spanContext.SpanID)
		}
		curLevel = nextLevel
	}
	s.tracer.mu.Unlock()
	s.tracer.logger.Printf("SimpleTrace %s", strings.Join(parts, " "))
}

func (s *SimpleSpan) Context() SpanContext {
	return s.spanContext
}

func (s *SimpleSpan) SetOperationName(operationName string) Span {
	s.operationName = operationName
	return s
}

func (s *SimpleSpan) Tracer() Tracer {
	return s.tracer
}

func (s *SimpleSpan) isRoot() bool {
	return s.parentID == RootParentID
}

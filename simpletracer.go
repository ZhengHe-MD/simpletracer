package simpletracer

import (
	"intracer/api"
	"math/rand"
	"sync"
	"time"
)

type SimpleTracer struct {
	mu              sync.RWMutex
	randomNumber    func() uint64
	spanIDToSpan    map[uint64]*SimpleSpan
	parentIDToSpans map[uint64][]*SimpleSpan
	logger          Logger
}

func NewTracer() api.Tracer {
	// randomNumber
	seedGenerator := NewRand(time.Now().UnixNano())
	pool := sync.Pool{
		New: func() interface{} {
			return rand.NewSource(seedGenerator.Int63())
		},
	}

	randomNumber := func() uint64 {
		generator := pool.Get().(rand.Source)
		n := uint64(generator.Int63())
		pool.Put(generator)
		return n
	}

	return &SimpleTracer{
		randomNumber:    randomNumber,
		spanIDToSpan:    make(map[uint64]*SimpleSpan),
		parentIDToSpans: make(map[uint64][]*SimpleSpan),
		logger:          defaultLogger{},
	}
}

func (t *SimpleTracer) StartSpan(operationName string, opts ...api.StartSpanOption) api.Span {
	sso := api.StartSpanOptions{}
	for _, opt := range opts {
		opt.Apply(&sso)
	}
	span := newInSpan(t, operationName, sso)

	t.mu.Lock()
	t.spanIDToSpan[span.spanContext.SpanID] = span
	if span.parentID != RootParentID {
		t.parentIDToSpans[span.parentID] = append(
			t.parentIDToSpans[span.parentID], span)
	}
	t.mu.Unlock()
	return span
}

// randomID generates a random trace/span ID, using tracer.random() generator.
// It never returns 0
func (t *SimpleTracer) randomID() uint64 {
	val := t.randomNumber()
	for val == 0 {
		val = t.randomNumber()
	}
	return val
}

func init() {
	api.SetGlobalTracer(NewTracer())
}
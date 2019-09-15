package simpletracer

import (
	"context"
	"testing"
)

func funcA(ctx context.Context) {
	span, ctx := StartSpanFromContext(ctx, "A")
	defer span.Finish()
	funcB(ctx)
	funcC(ctx)
}

func funcB(ctx context.Context) {
	span, ctx := StartSpanFromContext(ctx, "B")
	defer span.Finish()
	funcD(ctx)
}

func funcC(ctx context.Context) {
	span, ctx := StartSpanFromContext(ctx, "C")
	defer span.Finish()
	funcD(ctx)
}

func funcD(ctx context.Context) {
	span, ctx := StartSpanFromContext(ctx, "D")
	defer span.Finish()
}

func TestFuncA(t *testing.T) {
	funcA(context.Background())
}

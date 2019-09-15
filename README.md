warn: this project is under development, not production ready

# Simple Tracer

simple tracer is designed for trace inside standalone service. The apis are intentionally taken after opentracing specs.

It's likely that you want to trace functions inside a service instance. The function calls form a DAG, and you want to know how much time each function call takes

```go
func funcA() {
	funcB()
	funcC()
}

In the case of distributed tracing, we can use the popular open-sourced projects like Jaeger, that implements opentracing specs. The trace data will be sampled and collected by some agents and collector, and indexed by a search engine and finally an Web UI for easy lookups.

In the case of standalone tracing, we usually s
func funcB() {
	funcD()
}

func funcC() {
	funcD()
}

func funcD() {}
```
imply log the trace, manually, in each function call, and search that logs in Kibana with several clicks. It would be better if we can log it once, in the root function call of the DAG, and search it once. That's what this project is about.

change the above codes to the following
```go
import "github.com/ZhengHe-MD/simpletracer"

func funcA(ctx context.Context) {
	span, ctx := simpletracer.StartSpanFromContext(ctx, "A")
	defer span.Finish()
	funcB(ctx)
	funcC(ctx)
}

func funcB(ctx context.Context) {
	span, ctx := simpletracer.StartSpanFromContext(ctx, "B")
	defer span.Finish()
	funcD(ctx)
}

func funcC(ctx context.Context) {
	span, ctx := simpletracer.StartSpanFromContext(ctx, "C")
	defer span.Finish()
	funcD(ctx)
}

func funcD(ctx context.Context) {
	span, ctx := simpletracer.StartSpanFromContext(ctx, "D")
	defer span.Finish()
}
``` 
execute the codes, you will find the following info printed by the logger: `SimpleTrace A:12.505µs B:5.721µs C:1.427µs D:630ns D:419ns`




 
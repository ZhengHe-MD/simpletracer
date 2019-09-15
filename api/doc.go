// all codes of this package originate from opentracing/opentracing-go project
// we copy them because opentracing-go uses a singleton global tracer, if we simply
// implements a tracer for our purpose, application will not be able to use the
// distributed implementation, like Jaeger, at the same time.
package api

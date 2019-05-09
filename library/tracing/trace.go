package tracing

import "github.com/opentracing/opentracing-go"

var (
	_tracer opentracing.Tracer = &nooptracer{}
)

// SetGlobalTracer SetGlobalTracer
func SetGlobalTracer(tracer opentracing.Tracer) {
	_tracer = tracer
}

func StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	return _tracer.StartSpan(operationName, opts...)
}
func Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return _tracer.Inject(sm, format, carrier)
}
func Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	return _tracer.Extract(format, carrier)
}

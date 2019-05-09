package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

var (
	_ opentracing.Tracer = &nooptracer{}
)

type nooptracer struct{}

func (p *nooptracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	return &noopspan{}
}
func (p *nooptracer) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return nil
}
func (p *nooptracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	return &noopspancontext{}, nil
}

type noopspan struct{}

func (p *noopspan) Finish() {
}
func (p *noopspan) FinishWithOptions(opts opentracing.FinishOptions) {
}
func (p *noopspan) Context() opentracing.SpanContext {
	return &noopspancontext{}
}
func (p *noopspan) SetOperationName(operationName string) opentracing.Span {
	return p
}
func (p *noopspan) SetTag(key string, value interface{}) opentracing.Span {
	return p
}
func (p *noopspan) LogFields(fields ...log.Field) {
}
func (p *noopspan) LogKV(alternatingKeyValues ...interface{}) {
}
func (p *noopspan) SetBaggageItem(restrictedKey, value string) opentracing.Span {
	return p
}
func (p *noopspan) BaggageItem(restrictedKey string) string {
	return restrictedKey
}
func (p *noopspan) Tracer() opentracing.Tracer {
	return &nooptracer{}
}
func (p *noopspan) LogEvent(event string) {
}
func (p *noopspan) LogEventWithPayload(event string, payload interface{}) {
}
func (p *noopspan) Log(data opentracing.LogData) {
}

type noopspancontext struct {
}

func (p *noopspancontext) ForeachBaggageItem(handler func(k, v string) bool) {
}

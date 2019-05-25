package mars

import (
	"io"
	"net/http"
	"net/http/httptrace"
	"valerian/library/tracing"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

type contextKey int

const (
	keyTracer contextKey = iota
)

// Tracer holds tracing details for one HTTP request.
type ClientTracer struct {
	tr   opentracing.Tracer
	root opentracing.Span
	sp   opentracing.Span
	opts *clientOptions
}

type clientOptions struct {
	operationName      string
	componentName      string
	disableClientTrace bool
	spanObserver       func(span opentracing.Span, r *http.Request)
}

// ClientOption contols the behavior of TraceRequest.
type ClientOption func(*clientOptions)

// OperationName returns a ClientOption that sets the operation
// name for the client-side span.
func OperationName(operationName string) ClientOption {
	return func(options *clientOptions) {
		options.operationName = operationName
	}
}

// ComponentName returns a ClientOption that sets the component
// name for the client-side span.
func ComponentName(componentName string) ClientOption {
	return func(options *clientOptions) {
		options.componentName = componentName
	}
}

// ClientTrace returns a ClientOption that turns on or off
// extra instrumentation via httptrace.WithClientTrace.
func ClientTrace(enabled bool) ClientOption {
	return func(options *clientOptions) {
		options.disableClientTrace = !enabled
	}
}

// ClientSpanObserver returns a ClientOption that observes the span
// for the client-side span.
func ClientSpanObserver(f func(span opentracing.Span, r *http.Request)) ClientOption {
	return func(options *clientOptions) {
		options.spanObserver = f
	}
}

func (h *ClientTracer) start(req *http.Request) opentracing.Span {
	if h.root == nil {
		parent := opentracing.SpanFromContext(req.Context())
		var spanctx opentracing.SpanContext
		if parent != nil {
			spanctx = parent.Context()
		}
		operationName := h.opts.operationName
		if operationName == "" {
			operationName = "HTTP Client"
		}
		root := h.tr.StartSpan(operationName, opentracing.ChildOf(spanctx))
		h.root = root
	}

	ctx := h.root.Context()
	h.sp = h.tr.StartSpan("HTTP "+req.Method, opentracing.ChildOf(ctx))
	ext.SpanKindRPCClient.Set(h.sp)

	componentName := h.opts.componentName
	if componentName == "" {
		componentName = defaultComponentName
	}
	ext.Component.Set(h.sp, componentName)

	return h.sp
}

// Finish finishes the span of the traced request.
func (h *ClientTracer) Finish() {
	if h.root != nil {
		h.root.Finish()
	}
}

// Span returns the root span of the traced request. This function
// should only be called after the request has been executed.
func (h *ClientTracer) Span() opentracing.Span {
	return h.root
}

func (h *ClientTracer) clientTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn:              h.getConn,
		GotConn:              h.gotConn,
		PutIdleConn:          h.putIdleConn,
		GotFirstResponseByte: h.gotFirstResponseByte,
		Got100Continue:       h.got100Continue,
		DNSStart:             h.dnsStart,
		DNSDone:              h.dnsDone,
		ConnectStart:         h.connectStart,
		ConnectDone:          h.connectDone,
		WroteHeaders:         h.wroteHeaders,
		Wait100Continue:      h.wait100Continue,
		WroteRequest:         h.wroteRequest,
	}
}

func (h *ClientTracer) getConn(hostPort string) {
	ext.HTTPUrl.Set(h.sp, hostPort)
	h.sp.LogFields(log.String("event", "GetConn"))
}

func (h *ClientTracer) gotConn(info httptrace.GotConnInfo) {
	h.sp.SetTag("net/http.reused", info.Reused)
	h.sp.SetTag("net/http.was_idle", info.WasIdle)
	h.sp.LogFields(log.String("event", "GotConn"))
}

func (h *ClientTracer) putIdleConn(error) {
	h.sp.LogFields(log.String("event", "PutIdleConn"))
}

func (h *ClientTracer) gotFirstResponseByte() {
	h.sp.LogFields(log.String("event", "GotFirstResponseByte"))
}

func (h *ClientTracer) got100Continue() {
	h.sp.LogFields(log.String("event", "Got100Continue"))
}

func (h *ClientTracer) dnsStart(info httptrace.DNSStartInfo) {
	h.sp.LogFields(
		log.String("event", "DNSStart"),
		log.String("host", info.Host),
	)
}

func (h *ClientTracer) dnsDone(info httptrace.DNSDoneInfo) {
	fields := []log.Field{log.String("event", "DNSDone")}
	for _, addr := range info.Addrs {
		fields = append(fields, log.String("addr", addr.String()))
	}
	if info.Err != nil {
		fields = append(fields, log.Error(info.Err))
	}
	h.sp.LogFields(fields...)
}

func (h *ClientTracer) connectStart(network, addr string) {
	h.sp.LogFields(
		log.String("event", "ConnectStart"),
		log.String("network", network),
		log.String("addr", addr),
	)
}

func (h *ClientTracer) connectDone(network, addr string, err error) {
	if err != nil {
		h.sp.LogFields(
			log.String("message", "ConnectDone"),
			log.String("network", network),
			log.String("addr", addr),
			log.String("event", "error"),
			log.Error(err),
		)
	} else {
		h.sp.LogFields(
			log.String("event", "ConnectDone"),
			log.String("network", network),
			log.String("addr", addr),
		)
	}
}

func (h *ClientTracer) wroteHeaders() {
	h.sp.LogFields(log.String("event", "WroteHeaders"))
}

func (h *ClientTracer) wait100Continue() {
	h.sp.LogFields(log.String("event", "Wait100Continue"))
}

func (h *ClientTracer) wroteRequest(info httptrace.WroteRequestInfo) {
	if info.Err != nil {
		h.sp.LogFields(
			log.String("message", "WroteRequest"),
			log.String("event", "error"),
			log.Error(info.Err),
		)
		ext.Error.Set(h.sp, true)
	} else {
		h.sp.LogFields(log.String("event", "WroteRequest"))
	}
}

type closeTracker struct {
	io.ReadCloser
	sp opentracing.Span
}

func (c closeTracker) Close() error {
	err := c.ReadCloser.Close()
	c.sp.LogFields(log.String("event", "ClosedBody"))
	c.sp.Finish()
	return err
}

// NewTraceTracesport NewTraceTracesport
func NewTraceTracesport(rt http.RoundTripper, peerService string, internalTags ...tracing.Tag) *TraceTransport {
	return &TraceTransport{RoundTripper: rt, peerService: peerService, internalTags: internalTags}
}

// TraceTransport wraps a RoundTripper. If a request is being traced with
// Tracer, Transport will inject the current span into the headers,
// and set HTTP related tags on the span.
type TraceTransport struct {
	peerService  string
	internalTags []tracing.Tag
	// The actual RoundTripper to use for the request. A nil
	// RoundTripper defaults to http.DefaultTransport.
	http.RoundTripper
}

// TracerFromRequest retrieves the Tracer from the request. If the request does
// not have a Tracer it will return nil.
func TracerFromRequest(req *http.Request) *ClientTracer {
	tr, ok := req.Context().Value(keyTracer).(*ClientTracer)
	if !ok {
		return nil
	}
	return tr
}

// RoundTrip implements the RoundTripper interface
func (t *TraceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt := t.RoundTripper
	if rt == nil {
		rt = http.DefaultTransport
	}
	tracer := TracerFromRequest(req)
	if tracer == nil {
		return rt.RoundTrip(req)
	}

	tracer.start(req)

	ext.HTTPMethod.Set(tracer.sp, req.Method)
	ext.HTTPUrl.Set(tracer.sp, req.URL.String())
	ext.SpanKind.Set(tracer.sp, "client")
	ext.Component.Set(tracer.sp, defaultComponentName)
	tracer.opts.spanObserver(tracer.sp, req)
	if t.peerService != "" {
		ext.PeerService.Set(tracer.sp, t.peerService)
	}

	for _, v := range t.internalTags {
		tracer.sp.SetTag(v.Key, v.Value)
	}

	carrier := opentracing.HTTPHeadersCarrier(req.Header)
	tracer.sp.Tracer().Inject(tracer.sp.Context(), opentracing.HTTPHeaders, carrier)
	resp, err := rt.RoundTrip(req)

	if err != nil {
		tracer.sp.Finish()
		return resp, err
	}
	ext.HTTPStatusCode.Set(tracer.sp, uint16(resp.StatusCode))
	if resp.StatusCode >= http.StatusInternalServerError {
		ext.Error.Set(tracer.sp, true)
	}
	if req.Method == "HEAD" {
		tracer.sp.Finish()
	} else {
		resp.Body = closeTracker{resp.Body, tracer.sp}
	}
	return resp, nil

}

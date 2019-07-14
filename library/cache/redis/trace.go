package redis

import (
	"context"
	"fmt"
	"valerian/library/net/trace"
	"valerian/library/tracing"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	_traceComponentName = "library/cache/redis"
	_tracePeerService   = "redis"
	_traceSpanKind      = "client"
)

// var _internalTags = []trace.Tag{
// 	trace.TagString(trace.TagSpanKind, _traceSpanKind),
// 	trace.TagString(trace.TagComponent, _traceComponentName),
// 	trace.TagString(trace.TagPeerService, _tracePeerService),
// }

type traceConn struct {
	// tr for pipeline, if tr != nil meaning on pipeline
	span opentracing.Span
	ctx  context.Context
	// connTag include e.g. ip,port
	Addr string

	// origin redis conn
	Conn
	pending int
}

func (t *traceConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	root := opentracing.SpanFromContext(t.ctx)
	// NOTE: ignored empty commandName
	// current sdk will Do empty command after pipeline finished
	if root == nil || commandName == "" {
		return t.Conn.Do(commandName, args...)
	}

	parentCtx := root.Context()
	span := tracing.StartSpan("Redis:"+commandName, opentracing.ChildOf(parentCtx))
	defer span.Finish()

	ext.Component.Set(span, _traceComponentName)
	ext.SpanKind.Set(span, _traceSpanKind)
	ext.PeerService.Set(span, _tracePeerService)
	ext.PeerAddress.Set(span, t.Addr)

	statement := commandName
	if len(args) > 0 {
		statement += fmt.Sprintf(" %v", args[0])
	}
	span.SetTag("db.statement", statement)
	reply, err = t.Conn.Do(commandName, args...)
	return
}

func (t *traceConn) Send(commandName string, args ...interface{}) error {
	t.pending++
	root, ok := trace.FromContext(t.ctx)
	if !ok {
		return t.Conn.Send(commandName, args...)
	}
	if t.span == nil {
		t.span = root.Fork("", "Redis:Pipeline")
		t.span.SetTag(_internalTags...)
		t.span.SetTag(t.connTags...)
	}

	statement := commandName
	if len(args) > 0 {
		statement += fmt.Sprintf(" %v", args[0])
	}
	t.tr.SetLog(
		trace.Log(trace.LogEvent, "Send"),
		trace.Log("db.statement", statement),
	)
	err := t.Conn.Send(commandName, args...)
	if err != nil {
		t.tr.SetTag(trace.TagBool(trace.TagError, true))
		t.tr.SetLog(
			trace.Log(trace.LogEvent, "Send Fail"),
			trace.Log(trace.LogMessage, err.Error()),
		)
	}
	return err
}

func (t *traceConn) Flush() error {
	if t.tr == nil {
		return t.Conn.Flush()
	}
	t.tr.SetLog(trace.Log(trace.LogEvent, "Flush"))
	err := t.Conn.Flush()
	if err != nil {
		t.tr.SetTag(trace.TagBool(trace.TagError, true))
		t.tr.SetLog(
			trace.Log(trace.LogEvent, "Flush Fail"),
			trace.Log(trace.LogMessage, err.Error()),
		)
	}
	return err
}

func (t *traceConn) Receive() (reply interface{}, err error) {
	if t.tr == nil {
		return t.Conn.Receive()
	}
	t.tr.SetLog(trace.Log(trace.LogEvent, "Receive"))
	reply, err = t.Conn.Receive()
	if err != nil {
		t.tr.SetTag(trace.TagBool(trace.TagError, true))
		t.tr.SetLog(
			trace.Log(trace.LogEvent, "Receive Fail"),
			trace.Log(trace.LogMessage, err.Error()),
		)
	}
	if t.pending > 0 {
		t.pending--
	}
	if t.pending == 0 {
		t.tr.Finish(nil)
		t.tr = nil
	}
	return reply, err
}

func (t *traceConn) WithContext(ctx context.Context) Conn {
	t.ctx = ctx
	return t
}

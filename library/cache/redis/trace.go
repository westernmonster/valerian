package redis

import (
	"context"
	"fmt"
	"valerian/library/tracing"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

const (
	_traceComponentName = "library/cache/redis"
	_tracePeerService   = "redis"
	_traceSpanKind      = "client"
)

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
	root := opentracing.SpanFromContext(t.ctx)
	if root == nil {
		return t.Conn.Send(commandName, args...)
	}

	if t.span == nil {
		parentCtx := root.Context()
		t.span = tracing.StartSpan("Redis:Pipeline", opentracing.ChildOf(parentCtx))
		ext.Component.Set(t.span, _traceComponentName)
		ext.SpanKind.Set(t.span, _traceSpanKind)
		ext.PeerService.Set(t.span, _tracePeerService)
		ext.PeerAddress.Set(t.span, t.Addr)
	}

	statement := commandName
	if len(args) > 0 {
		statement += fmt.Sprintf(" %v", args[0])
	}
	t.span.LogFields(
		log.String("event", "Send"),
		log.String("db.statement", statement),
	)
	err := t.Conn.Send(commandName, args...)
	if err != nil {
		t.span.SetTag("error", true)
		t.span.LogFields(
			log.String("event", "Send Fail"),
			log.String("message", err.Error()),
		)
	}
	return err
}

func (t *traceConn) Flush() error {
	if t.span == nil {
		return t.Conn.Flush()
	}
	t.span.LogFields(log.String("event", "Flush"))
	err := t.Conn.Flush()
	if err != nil {
		t.span.SetTag("error", true)
		t.span.LogFields(
			log.String("event", "Flush Fail"),
			log.String("message", err.Error()),
		)
	}
	return err
}

func (t *traceConn) Receive() (reply interface{}, err error) {
	if t.span == nil {
		return t.Conn.Receive()
	}

	t.span.LogFields(log.String("event", "Receive"))
	reply, err = t.Conn.Receive()
	if err != nil {
		t.span.SetTag("error", true)
		t.span.LogFields(
			log.String("event", "Receive Fail"),
			log.String("message", err.Error()),
		)
	}
	if t.pending > 0 {
		t.pending--
	}
	if t.pending == 0 {
		t.span.Finish()
		t.span = nil
	}
	return reply, err
}

func (t *traceConn) WithContext(ctx context.Context) Conn {
	t.ctx = ctx
	return t
}

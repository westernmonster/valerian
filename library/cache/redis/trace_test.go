package redis

import (
	"context"
	"testing"
	"valerian/library/tracing"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

type mockConn struct{}

func (c *mockConn) Close() error { return nil }
func (c *mockConn) Err() error   { return nil }
func (c *mockConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return nil, nil
}
func (c *mockConn) Send(commandName string, args ...interface{}) error { return nil }
func (c *mockConn) Flush() error                                       { return nil }
func (c *mockConn) Receive() (reply interface{}, err error)            { return nil, nil }
func (c *mockConn) WithContext(context.Context) Conn                   { return c }

func TestTraceDo(t *testing.T) {
	tracer := mocktracer.New()
	tracing.SetGlobalTracer(tracer)
	tr := tracing.StartSpan("root")

	ctx := context.Background()
	parentCtx := opentracing.ContextWithSpan(ctx, tr)

	tc := &traceConn{Conn: &mockConn{}}
	conn := tc.WithContext(parentCtx)

	conn.Do("GET", "test")

	finishedSpans := tracer.FinishedSpans()
	assert.True(t, len(finishedSpans) == 1)
	assert.Equal(t, "Redis:GET", finishedSpans[0].OperationName)
	assert.NotEmpty(t, finishedSpans[0].Tags())
}

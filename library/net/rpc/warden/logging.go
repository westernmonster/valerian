package warden

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"

	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/metadata"
	"valerian/library/stat"
)

var (
	statsClient = stat.RPCClient
	statsServer = stat.RPCServer
)

func logFn(code int, dt time.Duration, fields []zapcore.Field) {
	switch {
	case code < 0:
		log.Error("", fields...)
		return
	case dt >= time.Millisecond*500:
		log.Warn("", fields...)
		return
	case code > 0:
		log.Warn("", fields...)
		return
	}
	log.Info("", fields...)
}

// clientLogging warden grpc logging
func clientLogging() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		var peerInfo peer.Peer
		opts = append(opts, grpc.Peer(&peerInfo))

		// invoker requests
		err := invoker(ctx, method, req, reply, cc, opts...)

		// after request
		code := ecode.Cause(err).Code()
		duration := time.Since(startTime)
		// monitor
		statsClient.Timing(method, int64(duration/time.Millisecond))
		statsClient.Incr(method, strconv.Itoa(code))

		var ip string
		if peerInfo.Addr != nil {
			ip = peerInfo.Addr.String()
		}
		logFields := []zapcore.Field{
			zap.String("ip", ip),
			zap.String("path", method),
			zap.Int("ret", code),
			// TODO: it will panic if someone remove String method from protobuf message struct that auto generate from protoc.
			zap.String("args", req.(fmt.Stringer).String()),
			zap.Float64("ts", duration.Seconds()),
			zap.String("source", "grpc-access-log"),
		}
		if err != nil {
			logFields = append(logFields, zap.String("error", err.Error()), zap.String("stack", fmt.Sprintf("%+v", err)))
		}
		logFn(code, duration, logFields)
		return err
	}
}

// serverLogging warden grpc logging
func serverLogging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		caller := metadata.String(ctx, metadata.Caller)
		var remoteIP string
		if peerInfo, ok := peer.FromContext(ctx); ok {
			remoteIP = peerInfo.Addr.String()
		}
		var quota float64
		if deadline, ok := ctx.Deadline(); ok {
			quota = time.Until(deadline).Seconds()
		}

		// call server handler
		resp, err := handler(ctx, req)

		// after server response
		code := ecode.Cause(err).Code()
		duration := time.Since(startTime)

		// monitor
		statsServer.Timing(caller, int64(duration/time.Millisecond), info.FullMethod)
		statsServer.Incr(caller, info.FullMethod, strconv.Itoa(code))
		logFields := []zapcore.Field{
			zap.String("user", caller),
			zap.String("ip", remoteIP),
			zap.String("path", info.FullMethod),
			zap.Int("ret", code),
			// TODO: it will panic if someone remove String method from protobuf message struct that auto generate from protoc.
			zap.String("args", req.(fmt.Stringer).String()),
			zap.Float64("ts", duration.Seconds()),
			zap.Float64("timeout_quota", quota),
			zap.String("source", "grpc-access-log"),
		}
		if err != nil {
			logFields = append(logFields, zap.String("error", err.Error()), zap.String("stack", fmt.Sprintf("%+v", err)))
		}
		logFn(code, duration, logFields)
		return resp, err
	}
}

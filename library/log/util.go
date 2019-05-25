package log

import (
	"context"
	"runtime"
	"strconv"
	"sync"
	"valerian/library/conf/env"
	"valerian/library/net/metadata"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var fm sync.Map

// funcName get func name.
func funcName(skip int) (name string) {
	if pc, _, lineNo, ok := runtime.Caller(skip); ok {
		if v, ok := fm.Load(pc); ok {
			name = v.(string)
		} else {
			name = runtime.FuncForPC(pc).Name() + ":" + strconv.FormatInt(int64(lineNo), 10)
			fm.Store(pc, name)
		}
	}
	return
}

func addExtraField(ctx context.Context, fields map[string]interface{}) {
	if caller := metadata.String(ctx, metadata.Caller); caller != "" {
		fields[_caller] = caller
	}
	if color := metadata.String(ctx, metadata.Color); color != "" {
		fields[_color] = color
	}
	if cluster := metadata.String(ctx, metadata.Cluster); cluster != "" {
		fields[_cluster] = cluster
	}
	fields[_deplyEnv] = env.DeployEnv
	fields[_zone] = env.Zone
	fields[_appID] = env.AppID
	fields[_instanceID] = env.Hostname
	if metadata.Bool(ctx, metadata.Mirror) {
		fields[_mirror] = true
	}
}

func appendFields() []zapcore.Field {
	fields := make([]zapcore.Field, 0)
	fields = append(fields, zap.String(_source, funcName(4)))

	return fields
}

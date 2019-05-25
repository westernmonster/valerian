package log

import (
	"fmt"
	"os"
	"time"
	"valerian/library/conf/env"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Family string
	Host   string

	// stdout
	Stdout bool

	Filter []string
}

var (
	l *zap.Logger
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func init() {
	host, _ := os.Hostname()
	l, _ = zap.NewDevelopment(
		zap.AddCaller(),
		zap.Fields(zap.String(_appID, env.AppID)),
		zap.Fields(zap.String(_instanceID, host)),
	)
}

func ZapLogger() *zap.Logger {
	return l
}

func Init(conf *Config) {
	atom := zap.NewAtomicLevelAt(zap.InfoLevel)
	config := zap.Config{
		Level:            atom,   // 日志级别
		Development:      true,   // 开发模式，堆栈跟踪
		Encoding:         "json", // 输出格式 console 或 json
		DisableCaller:    false,
		EncoderConfig:    NewEncoderConfig(),                                  // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout"},                                  // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}

	l = logger
}

// Info logs an info msg with fields
func Info(msg string, fields ...zapcore.Field) {
	fields = append(fields, appendFields()...)
	l.Info(msg, fields...)
}

// Error logs an error msg with fields
func Error(msg string, fields ...zapcore.Field) {
	l.Error(msg, fields...)
}

// Fatal logs a fatal error msg with fields
func Fatal(msg string, fields ...zapcore.Field) {
	l.Fatal(msg, fields...)
}

// Fatal logs a fatal error msg with fields
func Warn(msg string, fields ...zapcore.Field) {
	l.Warn(msg, fields...)
}

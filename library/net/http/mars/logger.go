package mars

import (
	"fmt"
	"strconv"
	"time"
	"valerian/library/ecode"
	"valerian/library/log"
	"valerian/library/net/metadata"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is logger  middleware
func Logger() HandlerFunc {
	const noUser = "no_user"
	return func(c *Context) {
		now := time.Now()
		ip := metadata.String(c, metadata.RemoteIP)
		req := c.Request
		path := req.URL.Path
		params := req.Form
		var quota float64
		if deadline, ok := c.Context.Deadline(); ok {
			quota = time.Until(deadline).Seconds()
		}

		c.Next()

		aid, _ := c.Get("aid")
		err := c.Error
		cerr := ecode.Cause(err)
		dt := time.Since(now)
		caller := metadata.String(c, metadata.Caller)
		if caller == "" {
			caller = noUser
		}

		stats.Incr(caller, path[1:], strconv.FormatInt(int64(cerr.Code()), 10))
		stats.Timing(caller, int64(dt/time.Millisecond), path[1:])

		errmsg := ""
		isSlow := dt >= (time.Millisecond * 500)

		fields := []zapcore.Field{
			zap.String("method", req.Method),
			zap.Any("aid", aid),
			zap.String("ip", ip),
			zap.String("user", caller),
			zap.String("path", path),
			zap.String("params", params.Encode()),
			zap.Int("ret", cerr.Code()),
			zap.String("msg", cerr.Message()),
			zap.String("stack", fmt.Sprintf("%+v", err)),
			zap.Float64("timeout_quota", quota),
			zap.Float64("ts", dt.Seconds()),
			zap.String("source", "http-access-log"),
		}

		if err != nil {
			errmsg = err.Error()
			fields = append(fields, zap.String("err", errmsg))
			if cerr.Code() > 0 {
				log.For(c).Warn("http", fields...)
				return
			}
			log.For(c).Error("http", fields...)
			return

		} else {
			if isSlow {
				log.For(c).Warn("http", fields...)
				return
			}
		}

		if path != "/monitor/ping" {
			log.For(c).Info("http", fields...)
		}
	}
}

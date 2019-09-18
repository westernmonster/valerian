// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tracing

import (
	"flag"
	"fmt"
	"os"
	"time"

	"valerian/library/conf/dsn"
	"valerian/library/conf/env"
	"valerian/library/log"
	xtime "valerian/library/time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
	jprom "github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
)

var _traceDSN = "127.0.0.1:6831"

func init() {
	if v := os.Getenv("TRACE"); v != "" {
		_traceDSN = v
	}
	flag.StringVar(&_traceDSN, "trace", _traceDSN, "trace report dsn, or use TRACE env.")
}

type Config struct {
	Address string
	// Report timeout
	Timeout xtime.Duration `dsn:"query.timeout,200ms"`
	// DisableSample
	DisableSample bool `dsn:"query.disable_sample"`
	// probabilitySampling
	Probability float32 `dsn:"-"`
}

func serviceNameFromEnv() string {
	return env.AppID
}

func isUATEnv() float64 {
	if env.DeployEnv == env.DeployEnvUat {
		return 0
	}

	return 1
}

func parseDSN(rawdsn string) (*Config, error) {
	d, err := dsn.Parse(rawdsn)
	if err != nil {
		return nil, errors.Wrapf(err, "trace: invalid dsn: %s", rawdsn)
	}

	cfg := new(Config)
	if _, err = d.Bind(cfg); err != nil {
		return nil, errors.Wrapf(err, "trace: invalid dsn: %s", rawdsn)
	}
	return cfg, nil
}

// Init creates a new instance of Jaeger tracer.
func Init(c *Config) opentracing.Tracer {
	serviceName := serviceNameFromEnv()
	metricsFactory := jprom.New().Namespace(metrics.NSOptions{Name: serviceName})

	if c == nil {
		// cfg, err := parseDSN(_traceDSN)
		// if err != nil {
		// 	panic(fmt.Errorf("parse trace dsn error: %s", err))
		// }

		c = &Config{
			Address: _traceDSN,
		}
	}

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
			// Param: isUATEnv(),
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  c.Address,
		},
	}

	logger := log.NewFactory()
	tracer, _, err := cfg.New(
		serviceName,
		config.Logger(jaegerLoggerAdapter{logger.Bg()}),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		logger.Bg().Fatal("cannot initialize Jaeger Tracer", zap.Error(err))
	}

	SetGlobalTracer(tracer)

	return tracer

}

type jaegerLoggerAdapter struct {
	logger log.Logger
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(msg)
}

func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(msg, args...))
}

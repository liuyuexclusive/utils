package trace

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/yuexclusive/utils/config"
)

func Tracer() (opentracing.Tracer, io.Closer, error) {
	cfg := config.MustGet()
	jcfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		ServiceName: cfg.Name,
	}

	report := jaegercfg.ReporterConfig{
		LogSpans:           true,
		LocalAgentHostPort: cfg.JaegerAddress,
		QueueSize:          1000,
	}

	reporter, _ := report.NewReporter(cfg.Name, jaeger.NewNullMetrics(), jaeger.NullLogger)
	return jcfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)
}

package tracex

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type Option func() attribute.KeyValue

func SetTracerProvider(url string, opts ...Option) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	attrs := make([]attribute.KeyValue, 0, len(opts))
	for _, opt := range opts {
		attrs = append(attrs, opt())
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(attrs...)),
	)
	otel.SetTracerProvider(tp)
	return err
}

func WithTracerAttrString(k, v string) Option {
	return func() attribute.KeyValue {
		return attribute.String(k, v)
	}
}

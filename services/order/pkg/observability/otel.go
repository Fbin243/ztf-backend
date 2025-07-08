package observability

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	stdoutlog "go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	otel_log "go.opentelemetry.io/otel/sdk/log"

	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func Init(serviceName string) error {
	// Trace exporter (stdout)
	// traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	// if err != nil {
	// 	return err
	// }

	// OLTP exporter (jaeger)
	ctx := context.Background()
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Gracefully shutdown the tracer provider on exit
	go func() {
		<-context.Background().Done()
		_ = tp.Shutdown(context.Background())
	}()

	// Log exporter (stdout)
	logExporter, err := stdoutlog.New(stdoutlog.WithPrettyPrint())
	if err != nil {
		return err
	}
	logProcessor := otel_log.NewBatchProcessor(logExporter)
	logProvider := otel_log.NewLoggerProvider(
		otel_log.WithProcessor(logProcessor),
		otel_log.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceName(serviceName),
			),
		))
	logger := otelslog.NewLogger(serviceName, otelslog.WithLoggerProvider(logProvider))

	slog.SetDefault(logger)

	return nil
}

func LogTraceAndSpanID(ctx context.Context, name string) {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		slog.Info("❌ No valid span context found")
		return
	}
	slog.Info(
		name,
		"TraceID",
		spanCtx.TraceID().String(),
		"SpanID",
		spanCtx.SpanID().String(),
	)
}

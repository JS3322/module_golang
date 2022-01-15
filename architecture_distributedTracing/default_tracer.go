package main

import "github.com/opentracing/opentracing-go"
import ".../some_tracing_impl"

func main() {
	opentracing.SetGlobalTracer(
		// tracing impl specific:
		some_tracing_impl.New(...),
	)
	...
}

func xyz(ctx context.Context, ...) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "operation_name")
	defer span.Finish()
	span.LogFields(
		log.String("event", "soft error"),
		log.String("type", "cache timeout"),
		log.Int("waited.millis", 1500))
}
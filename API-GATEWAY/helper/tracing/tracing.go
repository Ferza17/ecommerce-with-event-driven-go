package tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func StartSpanFromRpc(tracer opentracing.Tracer, operationName string) opentracing.Span {
	spanCtx, _ := extractTextMap(tracer, operationName)
	return tracer.StartSpan(operationName, ext.RPCServerOption(spanCtx))
}

func StartSpanFromContext(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, string(utils.APIGatewaySpanContextKey))
	span.SetOperationName(operationName)
	return span, ctx
}

func extractTextMap(tracer opentracing.Tracer, operationName string) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.TextMap,
		opentracing.TextMapCarrier{
			"operation_name": operationName,
		})
}

package tracing

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	spanContextKey = "api_gateway_ctx"
)

func StartSpanFromHttpRequest(tracer opentracing.Tracer, operationName string, r *http.Request) opentracing.Span {
	spanCtx, _ := extractHttpHeaders(tracer, r)
	return tracer.StartSpan(operationName, ext.RPCServerOption(spanCtx))
}

func StartSpanFromContext(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, spanContextKey)
	span.SetOperationName(operationName)
	return span, ctx
}

func extractHttpHeaders(tracer opentracing.Tracer, r *http.Request) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
}

func InjectToHTTPHeaders(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

package tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/Ferza17/event-driven-account-service/utils"
)

func StartSpanFromContext(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, string(utils.AccountSpanContextKey))
	span.SetOperationName(operationName)
	return span, ctx
}

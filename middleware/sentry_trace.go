package middleware

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SentryTraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hub := sentrygin.GetHubFromContext(ctx)
		traceCtx := sentry.SetHubOnContext(ctx, hub)
		span := sentry.StartSpan(
			traceCtx,
			fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.RequestURI),
			sentry.ContinueFromRequest(ctx.Request),
		)
		ctx.Set("parentSpan", span)
		defer span.Finish()
		ctx.Next()
	}
}

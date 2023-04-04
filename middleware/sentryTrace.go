package middleware

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// SentryTraceMiddleware создает middleware для добавления в события Sentry стектрейса.
func SentryTraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transaction := sentry.StartTransaction(
			ctx,
			fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.RequestURI),
			sentry.OpName("http"),
			sentry.ContinueFromRequest(ctx.Request),
		)
		defer transaction.Finish()
		ctx.Next()
	}
}

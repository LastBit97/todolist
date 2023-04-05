package middleware

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

const (
	transactionContextKey string = "transaction"
	httpOperation         string = "http"
)

// SentryTraceMiddleware создает middleware для добавления в события Sentry стектрейса.
func SentryTraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hub := sentrygin.GetHubFromContext(ctx)
		tracingCtx := sentry.SetHubOnContext(ctx, hub)
		transaction := sentry.StartTransaction(
			tracingCtx,
			fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.RequestURI),
			sentry.OpName(httpOperation),
			sentry.ContinueFromRequest(ctx.Request),
		)
		defer transaction.Finish()
		ctx.Set(transactionContextKey, transaction)
		ctx.Next()
	}
}

package sentry

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/getsentry/sentry-go"
)

const (
	transactionContextKey string = "transaction"
	op                    string = "db"
)

type SentryDriver struct {
	dialect.Driver
}

func Trace(d dialect.Driver) dialect.Driver {
	drv := &SentryDriver{d}
	return drv
}

func (d *SentryDriver) Exec(ctx context.Context, query string, args, v any) error {
	transaction := ctx.Value(transactionContextKey).(*sentry.Span)
	span := transaction.StartChild(op)
	span.Description = query
	defer span.Finish()
	return d.Driver.Exec(ctx, query, args, v)
}

func (d *SentryDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	transaction := ctx.Value(transactionContextKey).(*sentry.Span)
	span := transaction.StartChild(op)
	span.Description = query
	defer span.Finish()
	return drv.ExecContext(ctx, query, args...)
}

func (d *SentryDriver) Query(ctx context.Context, query string, args, v any) error {
	transaction := ctx.Value(transactionContextKey).(*sentry.Span)
	span := transaction.StartChild(op)
	span.Description = query
	defer span.Finish()
	return d.Driver.Query(ctx, query, args, v)
}

func (d *SentryDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	transaction := ctx.Value(transactionContextKey).(*sentry.Span)
	span := transaction.StartChild(op)
	span.Description = query
	defer span.Finish()
	return drv.QueryContext(ctx, query, args...)
}

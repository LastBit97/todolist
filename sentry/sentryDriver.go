package sentry

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
)

type SentryDriver struct {
	dialect.Driver
}

func Trace(d dialect.Driver) dialect.Driver {
	drv := &SentryDriver{d}
	return drv
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *SentryDriver) Exec(ctx context.Context, query string, args, v any) error {

	return d.Driver.Exec(ctx, query, args, v)
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *SentryDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}

	return drv.ExecContext(ctx, query, args...)
}

// Query logs its params and calls the underlying driver Query method.
func (d *SentryDriver) Query(ctx context.Context, query string, args, v any) error {

	return d.Driver.Query(ctx, query, args, v)
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *SentryDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}

	return drv.QueryContext(ctx, query, args...)
}

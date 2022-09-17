package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	pgxV4 "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Query ...
type Query struct {
	Name     string
	QueryRaw string
}

// DB ...
type DB struct {
	pool *pgxpool.Pool
}

// GetContext is a wrapped method pgxscan.Get ...
func (db *DB) GetContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	return pgxscan.Get(ctx, db.pool, dest, q.QueryRaw, args...)
}

// SelectContext is a wrapped method pgxscan.Select ...
func (db *DB) SelectContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	return pgxscan.Select(ctx, db.pool, dest, q.QueryRaw, args...)
}

// ExecContext is a wrapped method pgxpool.Exec ...
func (db *DB) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext is a wrapped method pgxpool.Query ...
func (db *DB) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgxV4.Rows, error) {
	return db.pool.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext is a wrapped method pgxpool.QueryRow ...
func (db *DB) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgxV4.Row {
	return db.pool.QueryRow(ctx, q.QueryRaw, args...)
}

package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Client ...
type Client interface {
	Close() error
	DB() *DB
}

type client struct {
	db        *DB
	closeFunc context.CancelFunc
}

// NewClient ...
func NewClient(ctx context.Context, config *pgxpool.Config) (Client, error) {
	dbc, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	_, cancel := context.WithCancel(ctx)

	return &client{
		db:        &DB{pool: dbc},
		closeFunc: cancel,
	}, nil
}

func (c *client) Close() error {
	if c != nil {
		if c.closeFunc != nil {
			c.closeFunc()
		}
		if c.db != nil {
			c.db.pool.Close()
		}
	}

	return nil
}

func (c *client) DB() *DB {
	return c.db
}

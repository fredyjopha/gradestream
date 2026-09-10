package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, dns string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dns)
	if err != nil {
		return nil, fmt.Errorf("connexion pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}
	return pool, nil
}

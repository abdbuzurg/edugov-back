package postgres

import (
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

//Store implements repositories.Store interface
//Has the SQLC generated queries for postgreSQL and has pgxpool for transactions
type Store struct {
	*sqlc.Queries
	pool *pgxpool.Pool
}

//Creates database connections with the postgreSQL
func NewPostgresDB(ctx context.Context, globalConfig *config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(globalConfig.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

  //setting up the config for connection pool
  config.MaxConns = int32(globalConfig.DBMaxOpenConns)
  config.MaxConnIdleTime = time.Duration(globalConfig.DBConnMaxLifetime)

  pool, err := pgxpool.New(ctx, config.ConnString())
  if err != nil {
    return nil, fmt.Errorf("failed to create pgxpool: %w", err)
  }

	// ping to varify database conneection is alive
	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}

//Creates instance of a Store for PostgreSQL
func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		Queries: sqlc.New(pool),
    pool: pool,
	}
}

//ExecTx executes a function within database transaction 
//Ensures fn are fully commited or rolled back together
//Uses pgxpool for transaction methods
func(s *Store) ExecTx(ctx context.Context, fn func(q *sqlc.Queries) error) error {
  tx, err := s.pool.Begin(ctx)
  if err != nil {
    return fmt.Errorf("failed to begin transaction: %w", err)
  }

  q := s.Queries.WithTx(tx)
  err = fn(q)
  if err != nil {
    if rbErr := tx.Rollback(ctx); rbErr != nil {
      return custom_errors.InternalServerError(fmt.Errorf("transaction failed: %w, rollback failed: %w", err, rbErr))
    }
    return err
  }

  return tx.Commit(ctx)
}

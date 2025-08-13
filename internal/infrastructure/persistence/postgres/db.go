package postgres

import (
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/persistence/postgres/sqlc"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

// Store implements repositories.Store interface
// Has the SQLC generated queries for postgreSQL and has pgxpool for transactions
type Store struct {
	*sqlc.Queries
	pool *pgxpool.Pool
}

// pgxLogger is an adapter that makes the standard Go logger compatible with the pgx v5 tracelog.Logger interface.
type pgxLogger struct {
	logger *log.Logger
}

// Log implements the tracelog.Logger interface.
func (l *pgxLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	// You can customize your log output here. This is a simple example.
	var logMsg strings.Builder
	logMsg.WriteString(msg)

	// Append data fields to the log message for more context.
	if len(data) > 0 {
		logMsg.WriteString(" |")
		for k, v := range data {
			logMsg.WriteString(fmt.Sprintf(" %s=%v", k, v))
		}
	}

	// Use the underlying standard logger to print the formatted message.
	l.logger.Println(logMsg.String())
}

// Creates database connections with the postgreSQL
func NewPostgresDB(ctx context.Context, globalConfig *config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(globalConfig.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	//setting up the config for connection pool
	// config.MaxConns = int32(globalConfig.DBMaxOpenConns)
	// config.MaxConnIdleTime = time.Duration(globalConfig.DBConnMaxLifetime)

	// loggerAdapter := &pgxLogger{
	// 	logger: log.New(os.Stdout, "PGX_DEBUG: ", log.LstdFlags),
	// }

	// tracer := &tracelog.TraceLog{
	// 	Logger:   loggerAdapter,
	// 	LogLevel: tracelog.LogLevelTrace,
	// }
	// // Assign the tracer to the connection config.
	// config.ConnConfig.Tracer = tracer

	pool, err := pgxpool.NewWithConfig(ctx, config)
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

// Creates instance of a Store for PostgreSQL
func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		Queries: sqlc.New(pool),
		pool:    pool,
	}
}

// ExecTx executes a function within database transaction
// Ensures fn are fully commited or rolled back together
// Uses pgxpool for transaction methods
func (s *Store) ExecTx(ctx context.Context, fn func(q *sqlc.Queries) error) error {
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
		return custom_errors.InternalServerError(fmt.Errorf("transaction failed: %w, rollback successful", err))
	}

	return tx.Commit(ctx)
}

// internal/infrastructure/persistence/postgres/store.go
package postgres

import (
	"context"
	"database/sql" // Use "database/sql" for generic SQL operations, or "github.com/jackc/pgx/v5/pgxpool" if using pgxpool directly
	"fmt"

	// Import your SQLC generated code
	"your-project/internal/infrastructure/persistence/postgres/sqlc"
)

// Store defines the interface for all database operations, including transaction management.
// It embeds the sqlc.Querier interface so that all generated SQLC query methods
// are directly available through the Store.
//
// Use cases that need to perform transactional operations will depend on this Store interface.
type Store interface {
	sqlc.Querier // Embeds all methods from sqlc.Queries (e.g., CreateUser, GetProductByID)

	// ExecTx executes a function within a new database transaction.
	// The function `fn` receives a *new* `sqlc.Queries` instance that is bound
	// to the active transaction. All database operations within `fn` using `qTx`
	// will be part of the same transaction.
	ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error
}

// SQLCStore implements the Store interface using the standard database/sql package.
// It holds the direct database connection and the base SQLC Queries object.
type SQLCStore struct {
	db *sql.DB // The underlying database connection pool
	*sqlc.Queries // The base SQLC Queries object, for non-transactional operations
}

// NewStore creates and returns a new SQLCStore.
// It initializes the base sqlc.Queries object with the provided database connection.
func NewStore(db *sql.DB) Store {
	return &SQLCStore{
		db:      db,
		Queries: sqlc.New(db), // Initialize with the direct DB connection for non-transactional use
	}
}

// ExecTx provides a transactional block for executing database operations.
// It handles the beginning, committing, and rolling back of a database transaction.
//
// The `fn` argument is a function that contains the actual database logic
// to be executed within the transaction. This function receives a `*sqlc.Queries`
// instance that is bound to the transaction, ensuring all operations within `fn`
// are part of the same atomic unit of work.
func (s *SQLCStore) ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	// Begin a new transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Create a new SQLC Queries instance that uses the transaction.
	// This is the key: all queries executed via `qTx` will be part of this transaction.
	qTx := sqlc.New(tx)

	// Execute the provided function `fn` with the transactional Queries.
	// This `fn` contains the business logic that needs to be atomic.
	err = fn(qTx)
	if err != nil {
		// If `fn` returns an error, rollback the transaction.
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction failed: %w; rollback failed: %v", err, rbErr)
		}
		return fmt.Errorf("transaction failed: %w", err)
	}

	// If `fn` completed without error, commit the transaction.
	return tx.Commit()
}

// NOTE: If you are using `pgx` (github.com/jackc/pgx/v5/pgxpool) instead of `database/sql`,
// you would adapt this file slightly. For example:
/*
import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool" // Use pgxpool for connection pooling
	"your-project/internal/infrastructure/persistence/postgres/sqlc"
)

type Store interface {
	sqlc.Querier
	ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error
}

type SQLCStore struct {
	connPool *pgxpool.Pool // Use pgxpool.Pool for the connection pool
	*sqlc.Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLCStore{
		connPool: connPool,
		Queries: sqlc.New(connPool), // pgxpool.Pool also implements sqlc.DBTX
	}
}

func (s *SQLCStore) ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	// Acquire a connection from the pool
	conn, err := s.connPool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection for transaction: %w", err)
	}
	defer conn.Release() // Release the connection back to the pool

	// Begin a new transaction on the acquired connection
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	qTx := sqlc.New(tx) // Create SQLC Queries bound to the transaction
	err = fn(qTx)

	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil { // Rollback also needs context for pgx
			return fmt.Errorf("transaction failed: %w; rollback failed: %v", err, rbErr)
		}
		return fmt.Errorf("transaction failed: %w", err)
	}

	return tx.Commit(ctx) // Commit also needs context for pgx
}
*/

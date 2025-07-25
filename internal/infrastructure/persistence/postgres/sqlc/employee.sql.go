// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: employee.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employees (
  unique_id
) VALUES(
  $1  
) RETURNING id, created_at, updated_at
`

type CreateEmployeeRow struct {
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateEmployee(ctx context.Context, uniqueID string) (CreateEmployeeRow, error) {
	row := q.db.QueryRow(ctx, createEmployee, uniqueID)
	var i CreateEmployeeRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM employees 
WHERE id = $1
`

func (q *Queries) DeleteEmployee(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteEmployee, id)
	return err
}

const getEmployeeByID = `-- name: GetEmployeeByID :one
SELECT id, unique_id, created_at, updated_at
FROM employees
WHERE id = $1
`

func (q *Queries) GetEmployeeByID(ctx context.Context, id int64) (Employee, error) {
	row := q.db.QueryRow(ctx, getEmployeeByID, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.UniqueID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEmployeeByUniqueIdentifier = `-- name: GetEmployeeByUniqueIdentifier :one
SELECT id, unique_id, created_at, updated_at
FROM employees
WHERE unique_id = $1
`

func (q *Queries) GetEmployeeByUniqueIdentifier(ctx context.Context, uniqueID string) (Employee, error) {
	row := q.db.QueryRow(ctx, getEmployeeByUniqueIdentifier, uniqueID)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.UniqueID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

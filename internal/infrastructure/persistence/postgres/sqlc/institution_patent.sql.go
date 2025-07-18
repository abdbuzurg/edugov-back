// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: institution_patent.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInstitutionPatent = `-- name: CreateInstitutionPatent :one
INSERT INTO institution_patents(
  institution_id,
  language_code,
  patent_title,
  discipline,
  description,
  implemented_in,
  link_to_patent_file
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING id, created_at, updated_at
`

type CreateInstitutionPatentParams struct {
	InstitutionID    int64  `json:"institution_id"`
	LanguageCode     string `json:"language_code"`
	PatentTitle      string `json:"patent_title"`
	Discipline       string `json:"discipline"`
	Description      string `json:"description"`
	ImplementedIn    string `json:"implemented_in"`
	LinkToPatentFile string `json:"link_to_patent_file"`
}

type CreateInstitutionPatentRow struct {
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateInstitutionPatent(ctx context.Context, arg CreateInstitutionPatentParams) (CreateInstitutionPatentRow, error) {
	row := q.db.QueryRow(ctx, createInstitutionPatent,
		arg.InstitutionID,
		arg.LanguageCode,
		arg.PatentTitle,
		arg.Discipline,
		arg.Description,
		arg.ImplementedIn,
		arg.LinkToPatentFile,
	)
	var i CreateInstitutionPatentRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

const deleteInstitutionPatent = `-- name: DeleteInstitutionPatent :exec
DELETE FROM institution_patents
WHERE id = $1
`

func (q *Queries) DeleteInstitutionPatent(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteInstitutionPatent, id)
	return err
}

const getInstitutionPatentByID = `-- name: GetInstitutionPatentByID :one
SELECT id, institution_id, language_code, patent_title, discipline, description, implemented_in, link_to_patent_file, created_at, updated_at
FROM institution_patents
WHERE id = $1
`

func (q *Queries) GetInstitutionPatentByID(ctx context.Context, id int64) (InstitutionPatent, error) {
	row := q.db.QueryRow(ctx, getInstitutionPatentByID, id)
	var i InstitutionPatent
	err := row.Scan(
		&i.ID,
		&i.InstitutionID,
		&i.LanguageCode,
		&i.PatentTitle,
		&i.Discipline,
		&i.Description,
		&i.ImplementedIn,
		&i.LinkToPatentFile,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getInstitutionPatentsByInstitutionIDAndLanguageCode = `-- name: GetInstitutionPatentsByInstitutionIDAndLanguageCode :many
SELECT id, institution_id, language_code, patent_title, discipline, description, implemented_in, link_to_patent_file, created_at, updated_at
FROM institution_patents
WHERE institution_id = $1 AND language_code = $2
`

type GetInstitutionPatentsByInstitutionIDAndLanguageCodeParams struct {
	InstitutionID int64  `json:"institution_id"`
	LanguageCode  string `json:"language_code"`
}

func (q *Queries) GetInstitutionPatentsByInstitutionIDAndLanguageCode(ctx context.Context, arg GetInstitutionPatentsByInstitutionIDAndLanguageCodeParams) ([]InstitutionPatent, error) {
	rows, err := q.db.Query(ctx, getInstitutionPatentsByInstitutionIDAndLanguageCode, arg.InstitutionID, arg.LanguageCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []InstitutionPatent{}
	for rows.Next() {
		var i InstitutionPatent
		if err := rows.Scan(
			&i.ID,
			&i.InstitutionID,
			&i.LanguageCode,
			&i.PatentTitle,
			&i.Discipline,
			&i.Description,
			&i.ImplementedIn,
			&i.LinkToPatentFile,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInstitutionPatent = `-- name: UpdateInstitutionPatent :one
UPDATE institution_patents
SET
  patent_title = COALESCE($1, patent_title),
  discipline = COALESCE($2, discipline),
  description = COALESCE($3, description),
  implemented_in = COALESCE($4, implemented_in),
  link_to_patent_file = COALESCE($5, link_to_patent_file),
  updated_at = now()
WHERE id = $6
RETURNING id, created_at, updated_at
`

type UpdateInstitutionPatentParams struct {
	PatentTitle      string `json:"patent_title"`
	Discipline       string `json:"discipline"`
	Description      string `json:"description"`
	ImplementedIn    string `json:"implemented_in"`
	LinkToPatentFile string `json:"link_to_patent_file"`
	ID               int64  `json:"id"`
}

type UpdateInstitutionPatentRow struct {
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) UpdateInstitutionPatent(ctx context.Context, arg UpdateInstitutionPatentParams) (UpdateInstitutionPatentRow, error) {
	row := q.db.QueryRow(ctx, updateInstitutionPatent,
		arg.PatentTitle,
		arg.Discipline,
		arg.Description,
		arg.ImplementedIn,
		arg.LinkToPatentFile,
		arg.ID,
	)
	var i UpdateInstitutionPatentRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

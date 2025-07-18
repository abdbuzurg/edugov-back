// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: institution_conference.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInstitutionConference = `-- name: CreateInstitutionConference :one
INSERT INTO institution_conferences (
  institution_id,
  language_code,
  conference_title,
  link,
  link_to_rinc,
  date_of_conference
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, created_at, updated_at
`

type CreateInstitutionConferenceParams struct {
	InstitutionID    int64       `json:"institution_id"`
	LanguageCode     string      `json:"language_code"`
	ConferenceTitle  string      `json:"conference_title"`
	Link             string      `json:"link"`
	LinkToRinc       pgtype.Text `json:"link_to_rinc"`
	DateOfConference pgtype.Date `json:"date_of_conference"`
}

type CreateInstitutionConferenceRow struct {
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateInstitutionConference(ctx context.Context, arg CreateInstitutionConferenceParams) (CreateInstitutionConferenceRow, error) {
	row := q.db.QueryRow(ctx, createInstitutionConference,
		arg.InstitutionID,
		arg.LanguageCode,
		arg.ConferenceTitle,
		arg.Link,
		arg.LinkToRinc,
		arg.DateOfConference,
	)
	var i CreateInstitutionConferenceRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

const deleteInstitutionConference = `-- name: DeleteInstitutionConference :exec
DELETE FROM institution_conferences
WHERE id = $1
`

func (q *Queries) DeleteInstitutionConference(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteInstitutionConference, id)
	return err
}

const getInstitutionConferenceByID = `-- name: GetInstitutionConferenceByID :one
SELECT id, institution_id, language_code, conference_title, link, link_to_rinc, date_of_conference, created_at, updated_at
FROM institution_conferences
WHERE id = $1
`

func (q *Queries) GetInstitutionConferenceByID(ctx context.Context, id int64) (InstitutionConference, error) {
	row := q.db.QueryRow(ctx, getInstitutionConferenceByID, id)
	var i InstitutionConference
	err := row.Scan(
		&i.ID,
		&i.InstitutionID,
		&i.LanguageCode,
		&i.ConferenceTitle,
		&i.Link,
		&i.LinkToRinc,
		&i.DateOfConference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getInstitutionConferencesByInstitutionIDAndLanguageCode = `-- name: GetInstitutionConferencesByInstitutionIDAndLanguageCode :many
SELECT id, institution_id, language_code, conference_title, link, link_to_rinc, date_of_conference, created_at, updated_at
FROM institution_conferences
WHERE institution_id = $1 AND language_code = $2
`

type GetInstitutionConferencesByInstitutionIDAndLanguageCodeParams struct {
	InstitutionID int64  `json:"institution_id"`
	LanguageCode  string `json:"language_code"`
}

func (q *Queries) GetInstitutionConferencesByInstitutionIDAndLanguageCode(ctx context.Context, arg GetInstitutionConferencesByInstitutionIDAndLanguageCodeParams) ([]InstitutionConference, error) {
	rows, err := q.db.Query(ctx, getInstitutionConferencesByInstitutionIDAndLanguageCode, arg.InstitutionID, arg.LanguageCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []InstitutionConference{}
	for rows.Next() {
		var i InstitutionConference
		if err := rows.Scan(
			&i.ID,
			&i.InstitutionID,
			&i.LanguageCode,
			&i.ConferenceTitle,
			&i.Link,
			&i.LinkToRinc,
			&i.DateOfConference,
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

const updateInstitutionConference = `-- name: UpdateInstitutionConference :one
UPDATE institution_conferences 
SET 
  conference_title = COALESCE($1, conference_title),
  link = COALESCE($2, link),
  link_to_rinc = COALESCE($3, link_to_rinc),
  date_of_conference = COALESCE($4, date_of_conference),
  updated_at = now()
WHERE id = $5
RETURNING id, created_at, updated_at
`

type UpdateInstitutionConferenceParams struct {
	ConferenceTitle  string      `json:"conference_title"`
	Link             string      `json:"link"`
	LinkToRinc       pgtype.Text `json:"link_to_rinc"`
	DateOfConference pgtype.Date `json:"date_of_conference"`
	ID               int64       `json:"id"`
}

type UpdateInstitutionConferenceRow struct {
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) UpdateInstitutionConference(ctx context.Context, arg UpdateInstitutionConferenceParams) (UpdateInstitutionConferenceRow, error) {
	row := q.db.QueryRow(ctx, updateInstitutionConference,
		arg.ConferenceTitle,
		arg.Link,
		arg.LinkToRinc,
		arg.DateOfConference,
		arg.ID,
	)
	var i UpdateInstitutionConferenceRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

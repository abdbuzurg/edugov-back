-- name: CreateInstitutionConference :one
INSERT INTO institution_conferences (
  institution_id,
  language_code,
  conference_title,
  link,
  link_to_rinc,
  date_of_conference
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionConference :one
UPDATE institution_conferences 
SET 
  conference_title = COALESCE($1, conference_title),
  link = COALESCE($2, link),
  link_to_rinc = COALESCE($3, link_to_rinc),
  date_of_conference = COALESCE($4, date_of_conference),
  updated_at = now()
WHERE id = $5
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionConference :exec
DELETE FROM institution_conferences
WHERE id = $1;

-- name: GetInstitutionConferenceByID :one
SELECT *
FROM institution_conferences
WHERE id = $1;

-- name: GetInstitutionConferencesByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_conferences
WHERE institution_id = $1 AND language_code = $2;

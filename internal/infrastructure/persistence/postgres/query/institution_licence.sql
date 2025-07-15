-- name: CreateInstitutionLicence :one
INSERT INTO institution_licences(
  institution_id,
  language_code,
  licence_title,
  licence_type,
  link_to_file,
  given_by,
  date_start,
  date_end
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionLicence :one
UPDATE institution_licences
SET
  licence_title = COALESCE($1, licence_title),
  licence_type = COALESCE($2, licence_type),
  link_to_file = COALESCE($3, link_to_file),
  given_by = COALESCE($4, given_by),
  date_start = COALESCE($5, date_start),
  date_end = COALESCE($6, date_end),
  updated_at = now()
WHERE id = $7
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionLicence :exec
DELETE FROM institution_licences
WHERE id = $1;

-- name: GetInstitutionLicenceByID :one
SELECT *
FROM institution_licences
WHERE id = $1;

-- name: GetInstitutionLicencesByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_licences
WHERE institution_id = $1 AND language_code = $2;

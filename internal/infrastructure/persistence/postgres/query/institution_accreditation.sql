-- name: CreateInstitutionAccreditation :one
INSERT INTO institution_accreditations(
  institution_id,
  language_code,
  accreditation_type,
  given_by
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionAccreditation :one
UPDATE institution_accreditations
SET 
  accreditation_type = COALESCE($1, accreditation_type),
  given_by = COALESCE($2, type),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionAccreditation :exec
DELETE FROM institution_accreditations
WHERE id = $1;

-- name: GetInstitutionAccreditationByID :one
SELECT *
FROM institution_accreditations
WHERE id = $1;

-- name: GetInstitutionAccreditationsByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_accreditations
WHERE institution_id = $1 and language_code = $2;

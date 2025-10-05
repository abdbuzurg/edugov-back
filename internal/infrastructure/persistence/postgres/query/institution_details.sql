-- name: CreateInstitutionDetails :one
INSERT INTO institution_details(
  institution_id,
  language_code,
  institution_type,
  institution_title_short,
  institution_title_long,
  legal_status,
  mission,
  founder,
  legal_address,
  factual_address,
  city
  )
VALUES (
 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, created_at, updated_at;

-- name: UpdateInsitutionDetails :one
UPDATE institution_details
SET 
  institution_type = COALESCE($1, institution_type),
  institution_title_long = COALESCE($2, institution_title_long),
  institution_title_short = COALESCE($3, institution_title_short),
  legal_status = COALESCE($4, legal_status),
  mission = COALESCE($5, mission),
  founder = COALESCE($6, founder),
  legal_address = COALESCE($7, legal_address),
  factual_address = COALESCE($8, factual_address),
  city = COALESCE($9, city),
  updated_at = now()
WHERE id = $10
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionDetails :exec
DELETE FROM institution_details 
WHERE id = $1;

-- name: GetInstitutionDetailsByID :one 
SELECT *
FROM institution_details
WHERE id = $1;

-- name: GetInstitutionDetailsByInstitutionIDAndLanguage :one
SELECT *
FROM institution_details
WHERE institution_id = $1 AND language_code = $2;

-- name: GetInstitutionNamesByLanguageCode :many
SELECT institution_title_long
FROM institution_details
WHERE language_code = $1;

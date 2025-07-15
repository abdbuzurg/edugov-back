-- name: CreateInstitutionDetails :one
INSERT INTO institution_details(
  institution_id,
  language_code,
  institution_type,
  institution_title,
  legal_status,
  mission,
  founder,
  legal_address,
  factual_address
  )
VALUES (
 $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, created_at, updated_at;

-- name: UpdateInsitutionDetails :one
UPDATE institution_details
SET 
  institution_type = COALESCE($1, institution_type),
  institution_title = COALESCE($2, institution_title),
  legal_status = COALESCE($3, legal_status),
  mission = COALESCE($4, mission),
  founder = COALESCE($5, founder),
  legal_address = COALESCE($6, legal_address),
  factual_address = COALESCE($7, factual_address),
  updated_at = now()
WHERE id = $8
RETURNING id, created_at, updated_at;

-- name: DeleteInsitutionDetails :exec
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

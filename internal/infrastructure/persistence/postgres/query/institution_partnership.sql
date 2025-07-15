-- name: CreateInstitutionPartnership :one
INSERT INTO institution_partnerships(
  institution_id,
  language_code,
  partner_name,
  partner_type,
  date_of_contract,
  link_to_partner,
  goal
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionPartnership :one
UPDATE institution_partnerships
SET
  partner_name = COALESCE($1, partner_name),
  partner_type = COALESCE($2, partner_type),
  date_of_contract = COALESCE($3, date_of_contract),
  link_to_partner = COALESCE($4, link_to_partner),
  goal = COALESCE($5, goal),
  updated_at = now()
WHERE id = $6
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionPartnership :exec
DELETE FROM institution_partnerships
WHERE id = $1;

-- name: GetInstitutionPartnershipByID :one
SELECT *
FROM institution_partnerships
WHERE id = $1;

-- name: GetInstitutionPartnershipsByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_partnerships
WHERE institution_id = $1 and language_code = $2;

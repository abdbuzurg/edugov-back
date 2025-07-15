-- name: CreateInstitutionProjectPartner :one
INSERT INTO institution_project_partners(
  institution_project_id,
  language_code,
  partner_name,
  partner_type,
  link_to_partner
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionProjectPartner :one
UPDATE institution_project_partners
SET
  partner_name = COALESCE($1, partner_name),
  partner_type = COALESCE($2, partner_type),
  link_to_partner = COALESCE($3, link_to_partner),
  updated_at = now()
WHERE id = $4
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionProjectPartner :exec
DELETE FROM institution_project_partners
WHERE id = $1;

-- name: GetInstitutionProjectPartnerByID :one
SELECT *
FROM institution_project_partners
WHERE id = $1;

-- name: GetInstitutionProjectPartnersByInstitutionProjectIDAndLanguageCode :many
SELECT *
FROM institution_project_partners
WHERE institution_project_id = $1 AND language_code = $2;

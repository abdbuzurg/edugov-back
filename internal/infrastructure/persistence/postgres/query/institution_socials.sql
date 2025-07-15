-- name: CreateInstitutionSocial :one
INSERT INTO institution_socials(
  institution_id,
  link_to_social,
  social_name
) VALUES (
  $1, $2, $3
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionSocial :one
UPDATE institution_socials
SET
  link_to_social = COALESCE($1, link_to_social),
  social_name = COALESCE($2, social_name),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionSocial :exec
DELETE FROM institution_socials
WHERE id = $1;

-- name: GetInstitutionSocialByID :one
SELECT *
FROM institution_socials
WHERE id = $1;

-- name: GetInstitutionSocialsByInstitutionID :many
SELECT *
FROM institution_socials
WHERE institution_id = $1;

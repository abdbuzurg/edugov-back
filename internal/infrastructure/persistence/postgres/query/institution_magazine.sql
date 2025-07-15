-- name: CreateInstitutionMagazine :one
INSERT INTO institution_magazines(
  institution_id,
  language_code,
  magazine_name,
  link,
  link_to_rinc
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionMagazine :one
UPDATE institution_magazines
SET 
  magazine_name = COALESCE($1, magazine_name),
  link = COALESCE($2, link),
  link_to_rinc = COALESCE($3, link_to_rinc),
  updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionMagazine :exec
DELETE FROM institution_magazines
WHERE id = $1;

-- name: GetInstitutionMagazineByID :one
SELECT *
FROM institution_magazines
WHERE id = $1;

-- name: GetInstitutionMagazinesByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_magazines
WHERE institution_id = $1 AND language_code = $2;

-- name: CreateInstitution :one
INSERT INTO institutions (
  year_of_establishment,
  email,
  fax,
  official_website,
  phone_number,
  mail_index
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitution :one
UPDATE institutions
SET 
  year_of_establishment = COALESCE($2, year_of_establishment),
  email = COALESCE($3, email),
  fax = COALESCE($4, fax),
  official_website = COALESCE($5, official_website),
  phone_number = COALESCE($6, phone_number),
  mail_index = COALESCE($7, mail_index),
  updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at;

-- name: DeleteInsitution :exec
DELETE FROM institutions
WHERE id = $1;

-- name: GetInstitutionByID :one
SELECT *
FROM institutions
WHERE 
  id = $1;

-- name: GetAllInstitutions :many
SELECT *
FROM institutions;

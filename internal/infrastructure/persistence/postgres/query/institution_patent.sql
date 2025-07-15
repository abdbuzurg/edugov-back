-- name: CreateInstitutionPatent :one
INSERT INTO institution_patents(
  institution_id,
  language_code,
  patent_title,
  discipline,
  description,
  implemented_in,
  link_to_patent_file
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionPatent :one
UPDATE institution_patents
SET
  patent_title = COALESCE($1, patent_title),
  discipline = COALESCE($2, discipline),
  description = COALESCE($3, description),
  implemented_in = COALESCE($4, implemented_in),
  link_to_patent_file = COALESCE($5, link_to_patent_file),
  updated_at = now()
WHERE id = $6
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionPatent :exec
DELETE FROM institution_patents
WHERE id = $1;

-- name: GetInstitutionPatentByID :one
SELECT *
FROM institution_patents
WHERE id = $1;

-- name: GetInstitutionPatentsByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_patents
WHERE institution_id = $1 AND language_code = $2;

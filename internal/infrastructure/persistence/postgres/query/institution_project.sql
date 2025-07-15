-- name: CreateInstitutionProject :one
INSERT INTO institution_projects(
  institution_id,
  language_code,
  project_type,
  project_title,
  date_start,
  date_end,
  fund,
  institution_role,
  coordinator
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionProject :one
UPDATE institution_projects
SET
  project_type = COALESCE($1, project_type),
  project_title = COALESCE($2, project_title),
  date_start = COALESCE($3, date_start),
  date_end = COALESCE($4, date_end),
  fund = COALESCE($5, fund),
  institution_role = COALESCE($6, institution_role),
  coordinator = COALESCE($7, coordinator),
  updated_at = now()
WHERE id = $8
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionProject :exec
DELETE FROM institution_projects
WHERE id = $1;

-- name: GetInstitutionProjectByID :one
SELECT *
FROM institution_projects
WHERE id = $1;

-- name: GetInstitutionProjectsByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_projects
WHERE institution_id = $1 AND language_code = $2;

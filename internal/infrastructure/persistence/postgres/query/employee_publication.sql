-- name: CreateEmployeePublication :one
INSERT INTO employee_publications(
  employee_id,
  language_code,
  publication_title,
  link_to_publication
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeePublication :one
UPDATE employee_publications 
SET 
  publication_title = COALESCE($1, publication_title),
  link_to_publication = COALESCE($2, link_to_publication),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeePublication :exec
DELETE FROM employee_publications
WHERE id = $1;

-- name: GetEmployeePublicationByID :one
SELECT *
FROM employee_publications
WHERE id = $1;

-- name: GetEmployeePublicationsByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_publications
WHERE employee_id = $1 AND language_code = $2;

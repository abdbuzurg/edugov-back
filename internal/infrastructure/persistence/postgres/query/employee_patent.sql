-- name: CreateEmployeePatent :one
INSERT INTO employee_patents(
  employee_id,
  language_code,
  patent_title,
  description
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeePatent :one
UPDATE employee_patents 
SET 
  patent_title = COALESCE($1, patent_title),
  description = COALESCE($2, description),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeePatent :exec
DELETE FROM employee_patents
WHERE id = $1;

-- name: GetEmployeePatentByID :one
SELECT *
FROM employee_patents
WHERE id = $1;

-- name: GetEmployeePatentsByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_patents
WHERE employee_id = $1 AND language_code = $2;

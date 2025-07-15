-- name: CreateEmployeeSocial :one
INSERT INTO employee_socials(
  employee_id,
  social_name,
  link_to_social
) VALUES (
  $1, $2, $3
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeSocial :one
UPDATE employee_socials 
SET 
  social_name = COALESCE($1, social_name),
  link_to_social = COALESCE($2, link_to_social),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeSocial :exec
DELETE FROM employee_socials
WHERE id = $1;

-- name: GetEmployeeSocialByID :one
SELECT *
FROM employee_socials
WHERE id = $1;

-- name: GetEmployeeSocialsByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_socials
WHERE employee_id = $1; 

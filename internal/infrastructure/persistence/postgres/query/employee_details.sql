-- name: CreateEmployeeDetails :one
INSERT INTO employee_details (
  employee_id,
  language_code,
  surname,
  name,
  middlename,
  is_employee_details_new
) VALUES(
  $1, $2, $3, $4, $5, $6
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeDetails :one
UPDATE employee_details
SET 
  surname = COALESCE($1, surname),
  name = COALESCE($2, name),
  middlename = COALESCE($3, middlename),
  is_employee_details_new = COALESCE($4, is_employee_details_new),
  updated_at = now()
WHERE id = $5
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeDetails :exec
DELETE FROM employee_details 
WHERE id = $1;

-- name: GetEmployeeDetailsByID :one
SELECT *
FROM employee_details
WHERE id = $1;

-- name: GetEmployeeDetailsByEmployeeID :many
SELECT *
FROM employee_details
WHERE employee_id = $1;

-- name: GetCurrentEmployeeDetailsByEmployeeIDAndLanguageCode :one
SELECT *
FROM employee_details
WHERE 
  employee_id = $1
  AND language_code = $2
  AND is_employee_details_new = 'true';
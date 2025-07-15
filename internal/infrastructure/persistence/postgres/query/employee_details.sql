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
  updated_at = now()
WHERE id = $4
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeDetails :exec
DELETE FROM employee_details 
WHERE id = $1;

-- name: GetEmployeeDetailsByID :one
SELECT *
FROM employee_details
WHERE id = $1;

-- name: GetEmployeeDetailsByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_details
WHERE employee_id = $1 and language_code = $2;

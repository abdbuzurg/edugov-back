-- name: CreateEmployee :one
INSERT INTO employees (
  unique_id
) VALUES(
  $1  
) RETURNING id, created_at, updated_at;

-- name: DeleteEmployee :exec
DELETE FROM employees 
WHERE id = $1;

-- name: GetEmployeeByID :one
SELECT *
FROM employees
WHERE id = $1;

-- name: GetEmployeeByUniqueIdentifier :one
SELECT *
FROM employees
WHERE unique_id = $1;

-- name: GetPersonnelPaginated :many
SELECT DISTINCT employees.id
FROM employees
INNER JOIN employee_details on employee_details.employee_id = employees.id
INNER JOIN employee_degrees on employee_degrees.employee_id = employees.id
INNER JOIN employee_work_experiences on employee_work_experiences.employee_id = employees.id
WHERE 
  -- MANDATORY FILTERS
  employee_details.language_code = sqlc.arg(language_code)
  AND employee_degrees.language_code = sqlc.arg(language_code)
  AND employee_work_experiences.language_code = sqlc.arg(language_code)
  -- OPTIONAL FILTERS
  AND (sqlc.narg(uid)::text IS NULL OR employees.uid = sqlc.narg(uid))
  AND (sqlc.narg(name)::text IS NULL OR employee_details.name = sqlc.narg(name))
  AND (sqlc.narg(surname)::text IS NULL OR employee_details.surname = sqlc.narg(surname))
  AND (sqlc.narg(middlename)::text IS NULL OR employee_details.middlename = sqlc.narg(uid))
  AND (sqlc.narg(speciality)::text IS NULL OR employees_degrees.speciality = sqlc.narg(speciality))
ORDER BY 
  employees.id ASC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg(page);
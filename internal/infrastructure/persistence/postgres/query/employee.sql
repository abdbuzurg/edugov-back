-- name: CreateEmployee :one
INSERT INTO employees (
  unique_id,
  user_id,
  gender,
  tin
) VALUES(
  $1, $2, $3, $4
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

-- name: GetEmployeeByUserID :one
SELECT * 
FROM employees
WHERE user_id = $1;

-- name: GetPersonnelPaginated :many
SELECT
	e.id,
	e.unique_id
FROM
	employees AS e
JOIN
	employee_details AS ed ON e.id = ed.employee_id
JOIN
	employee_degrees AS edeg ON e.id = edeg.employee_id
JOIN
	employee_work_experiences AS ewe ON e.id = ewe.employee_id
WHERE
	(sqlc.narg(uid)::text IS NULL OR e.unique_id ILIKE '%' || sqlc.narg(uid) || '%')
	AND ed.language_code = sqlc.arg(language_code)
	AND ed.is_employee_details_new = True
	AND (sqlc.narg(name)::text IS NULL OR ed.name ILIKE '%' || sqlc.narg(name) || '%')
	AND (sqlc.narg(surname)::text IS NULL OR ed.surname ILIKE '%' || sqlc.narg(surname) || '%')
	AND (sqlc.narg(middlename)::text IS NULL OR ed.middlename ILIKE '%' || sqlc.narg(middlename) || '%')
	AND edeg.language_code = sqlc.arg(language_code)
	AND (sqlc.narg(speciality)::text IS NULL OR edeg.speciality ILIKE '%' || sqlc.narg(speciality) || '%')
	AND ewe.language_code = sqlc.arg(language_code)
GROUP BY
	e.id,
	e.unique_id
ORDER BY
	e.id,
	e.unique_id
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg(page);

-- name: CountPersonnel :one
SELECT
	COUNT(DISTINCT e.id)
FROM
	employees AS e
JOIN
	employee_details AS ed ON e.id = ed.employee_id
JOIN
	employee_degrees AS edeg ON e.id = edeg.employee_id
JOIN
	employee_work_experiences AS ewe ON e.id = ewe.employee_id
WHERE
	(sqlc.narg(uid)::text IS NULL OR e.unique_id ILIKE '%' || sqlc.narg(uid) || '%')
	AND ed.language_code = sqlc.arg(language_code)
	AND ed.is_employee_details_new = True
	AND (sqlc.narg(name)::text IS NULL OR ed.name ILIKE '%' || sqlc.narg(name) || '%')
	AND (sqlc.narg(surname)::text IS NULL OR ed.surname ILIKE '%' || sqlc.narg(surname) || '%')
	AND (sqlc.narg(middlename)::text IS NULL OR ed.middlename ILIKE '%' || sqlc.narg(middlename) || '%')
	AND edeg.language_code = sqlc.arg(language_code)
	AND (sqlc.narg(speciality)::text IS NULL OR edeg.speciality ILIKE '%' || sqlc.narg(speciality) || '%')
	AND ewe.language_code = sqlc.arg(language_code);


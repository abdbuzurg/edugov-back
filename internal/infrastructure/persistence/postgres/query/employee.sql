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
	e.id as employee_id,
	e.unique_id as unique_id,
	ed.surname as surname,
	ed."name" as name,
	ed.middlename as middlename,
	latest_experience.workplace as currentWorkplace,
	latest_degree.degree_level as highestAcademicDegree,
	latest_degree.speciality as speciality
FROM employees AS e
JOIN employee_details AS ed ON e.id = ed.employee_id
JOIN (
	SELECT DISTINCT employee_id FROM employee_socials
) AS socials ON e.id = socials.employee_id
JOIN (
	SELECT DISTINCT on (employee_id)
		employee_id,
		workplace
	FROM employee_work_experiences
	WHERE employee_work_experiences.language_code = sqlc.arg(language_code)
	ORDER BY employee_work_experiences.employee_id, employee_work_experiences.date_start DESC, employee_work_experiences.on_going DESC
) AS latest_experience ON e.id = latest_experience.employee_id
JOIN (
	SELECT DISTINCT ON (employee_id)
		employee_id,
		degree_level,
		speciality
	FROM employee_degrees
	WHERE employee_degrees.language_code = sqlc.arg(language_code)
	ORDER BY employee_id, date_end desc
) AS latest_degree ON e.id = latest_degree.employee_id
WHERE
	(sqlc.narg(uid)::text IS NULL OR e.unique_id ILIKE '%' || sqlc.narg(uid) || '%')
	AND ed.language_code = sqlc.arg(language_code)
	AND ed.is_employee_details_new = true
	AND (sqlc.narg(name)::text IS NULL OR ed.name ILIKE '%' || sqlc.narg(name) || '%')
	AND (sqlc.narg(surname)::text IS NULL OR ed.surname ILIKE '%' || sqlc.narg(surname) || '%')
	AND (sqlc.narg(middlename)::text IS NULL OR ed.middlename ILIKE '%' || sqlc.narg(middlename) || '%')
ORDER BY e.id
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg(page);

-- name: CountPersonnel :one
SELECT COUNT(*) 
FROM (
	SELECT 
		e.id as employee_id
	FROM employees AS e
	JOIN employee_details AS ed ON e.id = ed.employee_id
	JOIN (
		SELECT DISTINCT employee_id FROM employee_socials
	) AS socials ON e.id = socials.employee_id
	JOIN (
		SELECT DISTINCT on (employee_id)
			employee_id,
			workplace
		FROM employee_work_experiences
		WHERE employee_work_experiences.language_code = sqlc.arg(language_code)
		ORDER by employee_work_experiences.employee_id, employee_work_experiences.date_end DESC NULLS FIRST
	) AS latest_experience ON e.id = latest_experience.employee_id
	JOIN (
		SELECT DISTINCT ON (employee_id)
			employee_id,
			degree_level,
			speciality
		FROM employee_degrees
		WHERE employee_degrees.language_code = sqlc.arg(language_code)
		ORDER BY employee_id, date_end desc
	) AS latest_degree ON e.id = latest_degree.employee_id
	WHERE
		(sqlc.narg(uid)::text IS NULL OR e.unique_id ILIKE '%' || sqlc.narg(uid) || '%')
		AND ed.language_code = sqlc.arg(language_code)
		AND ed.is_employee_details_new = true
		AND (sqlc.narg(name)::text IS NULL OR ed.name ILIKE '%' || sqlc.narg(name) || '%')
		AND (sqlc.narg(surname)::text IS NULL OR ed.surname ILIKE '%' || sqlc.narg(surname) || '%')
		AND (sqlc.narg(middlename)::text IS NULL OR ed.middlename ILIKE '%' || sqlc.narg(middlename) || '%')
) as final_result;

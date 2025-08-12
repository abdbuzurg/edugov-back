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
SELECT 
	employees.id,
	employees.unique_id
FROM 
	employees	
WHERE
	(sqlc.narg(uid)::text IS NULL or employees.unique_id = sqlc.narg(uid))
	AND EXISTS (
		SELECT 1
		FROM employee_details
		WHERE 
			employee_details.language_code = sqlc.arg(language_code)
			and employee_details.is_employee_details_new = True
		  AND (sqlc.narg(name)::text IS NULL OR employee_details.name = sqlc.narg(name))
		 	AND (sqlc.narg(surname)::text IS NULL OR employee_details.surname = sqlc.narg(surname))
		  AND (sqlc.narg(middlename)::text IS NULL OR employee_details.middlename = sqlc.narg(middlename))		
	  )
  	AND EXISTS (
  		SELECT 1
  		FROM employee_degrees
  		WHERE 
  			employee_degrees.employee_id = employees.id
  			and employee_degrees.language_code = sqlc.arg(language_code)
  			and (sqlc.narg(speciality)::text IS NULL OR employee_degrees.speciality = sqlc.narg(speciality))
  	)
  	AND EXISTS (
  		SELECT 1
  		FROM employee_work_experiences
  		WHERE 
  			employee_work_experiences.employee_id = employees.id
  			AND employee_work_experiences.language_code = sqlc.arg(language_code)
  	)
ORDER BY 
	employees.id,
	employees.unique_id
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg(page);

-- name: CountPersonnel :one
SELECT 
	COUNT(*)
FROM 
	employees	
WHERE
	(sqlc.narg(uid)::text IS NULL or employees.unique_id = sqlc.narg(uid))
	AND EXISTS (
		SELECT 1
		FROM employee_details
		WHERE 
			employee_details.language_code = sqlc.arg(language_code)
			and employee_details.is_employee_details_new = True
		  AND (sqlc.narg(name)::text IS NULL OR employee_details.name = sqlc.narg(name))
		 	AND (sqlc.narg(surname)::text IS NULL OR employee_details.surname = sqlc.narg(surname))
		  AND (sqlc.narg(middlename)::text IS NULL OR employee_details.middlename = sqlc.narg(middlename))		
	  )
  	AND EXISTS (
  		SELECT 1
  		FROM employee_degrees
  		WHERE 
  			employee_degrees.employee_id = employees.id
  			and employee_degrees.language_code = sqlc.arg(language_code)
  			and (sqlc.narg(speciality)::text IS NULL OR employee_degrees.speciality = sqlc.narg(speciality))
  	)
  	AND EXISTS (
  		SELECT 1
  		FROM employee_work_experiences
  		WHERE 
  			employee_work_experiences.employee_id = employees.id
  			AND employee_work_experiences.language_code = sqlc.arg(language_code)
  	);
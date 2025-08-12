-- name: CreateEmployeeDegree :one
INSERT INTO employee_degrees(
  employee_id,
  language_code,
  degree_level,
  university_name,
  speciality,
  date_start,
  date_end,
  given_by,
  date_degree_recieved
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeDegree :one
UPDATE employee_degrees 
SET 
  degree_level = COALESCE($1, degree_level),
  university_name = COALESCE($2, university_name),
  speciality = COALESCE($3, speciality),
  date_start = COALESCE($4, date_start),
  date_end = COALESCE($5, date_end),
  given_by = COALESCE($6, given_by),
  date_degree_recieved = COALESCE($7, date_degree_recieved),
  updated_at = now()
WHERE id = $8
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeDegree :exec
DELETE FROM employee_degrees
WHERE id = $1;

-- name: GetEmployeeDegreeByID :one
SELECT *
FROM employee_degrees
WHERE id = $1;

-- name: GetEmployeeDegreesByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_degrees
WHERE employee_id = $1 AND language_code = $2
ORDER BY employee_degrees.date_degree_recieved DESC;

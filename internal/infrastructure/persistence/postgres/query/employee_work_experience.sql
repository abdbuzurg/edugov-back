-- name: CreateEmployeeWorkExperience :one
INSERT INTO employee_work_experiences(
  employee_id,
  language_code,
  workplace,
  job_title,
  description,
  date_start,
  date_end,
  on_going
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeWorkExperience :one
UPDATE employee_work_experiences
SET 
  workplace = COALESCE($1, workplace),
  job_title = COALESCE($2, job_title),
  description = COALESCE($3, description),
  date_start = COALESCE($4, date_start),
  date_end = COALESCE($5, date_end),
  on_going = COALESCE($6, on_going),
  updated_at = now()
WHERE id = $7
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeWorkExperience :exec
DELETE FROM employee_work_experiences
WHERE id = $1;

-- name: GetEmployeeWorkExperienceByID :one
SELECT *
FROM employee_work_experiences
WHERE id = $1;

-- name: GetEmployeeWorkExperiencesByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_work_experiences
WHERE employee_id = $1 AND language_code = $2
ORDER BY employee_work_experiences.on_going DESC, employee_work_experiences.date_end DESC;
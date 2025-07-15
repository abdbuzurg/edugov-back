-- name: CreateEmployeeScientificAward :one
INSERT INTO employee_scientific_awards(
  employee_id,
  language_code,
  scientific_award_title,
  given_by
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeScientificAward :one
UPDATE employee_scientific_awards 
SET 
  scientific_award_title = COALESCE($1, scientific_award_title),
  given_by = COALESCE($2, given_by),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeScientificAward :exec
DELETE FROM employee_scientific_awards
WHERE id = $1;

-- name: GetEmployeeScientificAwardByID :one
SELECT *
FROM employee_scientific_awards
WHERE id = $1;

-- name: GetEmployeeScientificAwardsByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_scientific_awards
WHERE employee_id = $1 AND language_code = $2;

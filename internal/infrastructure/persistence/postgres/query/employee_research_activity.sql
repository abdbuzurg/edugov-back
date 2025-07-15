-- name: CreateEmployeeResearchActivity :one
INSERT INTO employee_research_activities(
  employee_id,
  language_code,
  research_activity_title,
  employee_role
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeResearchActivity :one
UPDATE employee_research_activities 
SET 
  research_activity_title = COALESCE($1, research_activity_title),
  employee_role = COALESCE($2, employee_role),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeResearchActivity :exec
DELETE FROM employee_research_activities
WHERE id = $1;

-- name: GetEmployeeResearchActivityByID :one
SELECT *
FROM employee_research_activities
WHERE id = $1;

-- name: GetEmployeeResearchActivitiesByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_research_activities
WHERE employee_id = $1 AND language_code = $2;

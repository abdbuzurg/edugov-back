-- name: CreateEmployeeParticipationInEvent :one
INSERT INTO employee_participation_in_events(
  employee_id,
  language_code,
  event_title,
  event_date
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeParticipationInEvent :one
UPDATE employee_participation_in_events 
SET 
  event_title = COALESCE($1, event_title),
  event_date = COALESCE($2, event_date),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeParticipationInEvent :exec
DELETE FROM employee_participation_in_events
WHERE id = $1;

-- name: GetEmployeeParticipationInEventByID :one
SELECT *
FROM employee_participation_in_events
WHERE id = $1;

-- name: GetEmployeeParticipationInEventsByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_participation_in_events
WHERE employee_id = $1 AND language_code = $2;

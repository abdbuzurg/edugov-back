-- name: CreateEmployeeMainResearchArea :one
INSERT INTO employee_main_research_areas(
  employee_id,
  language_code,
  area,
  discipline
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: CreateEmployeeMainResearchAreaKeyTopic :one
INSERT INTO employee_main_research_area_key_topics(
  employee_main_research_area_id, 
  key_topic_title
) VALUES (
  $1, $2
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeMainResearchArea :one
UPDATE employee_main_research_areas 
SET 
  area = COALESCE($1, area),
  discipline = COALESCE($2, discipline),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeMainResearchAreaKeyTopic :one
UPDATE employee_main_research_area_key_topics
SET
  key_topic_title = COALESCE($1, key_topic_title),
  updated_at = now()
WHERE id = $2
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeMainResearchArea :exec
DELETE FROM employee_main_research_areas
WHERE id = $1;

-- name: DeleteEmployeeMainResearchAreaKeyTopic :exec
DELETE FROM employee_main_research_area_key_topics
WHERE id = $1;

-- name: GetEmployeeMainResearchAreaByID :one
SELECT *
FROM employee_main_research_areas
WHERE id = $1;

-- name: GetEmployeeMainResearchAreasByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_main_research_areas
WHERE employee_id = $1 AND language_code = $2;

-- name: GetEmployeeMainResearchAreaKeyTopicByID :one
SELECT *
FROM employee_main_research_area_key_topics 
WHERE id = $1;

-- name: GetEmployeeMainResearchAreaKeyTopicsByEmployeeMainResearchAreaIDAndLanguageCode :many
SELECT *
FROM employee_main_research_area_key_topics
WHERE employee_main_research_area_id = $1;

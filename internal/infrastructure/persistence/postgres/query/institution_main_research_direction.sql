-- name: CreateInstitutionMainResearchDirection :one
INSERT INTO institution_main_research_directions(
  institution_id,
  language_code,
  research_direction_title,
  discipline,
  area_of_research
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionMainResearchDirection :one
UPDATE institution_main_research_directions
SET 
  research_direction_title = COALESCE($1, research_direction_title),
  discipline = COALESCE($2, discipline),
  area_of_research = COALESCE($3, area_of_research),
  updated_at = now()
WHERE id = $4
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionMainResearchDirection :exec
DELETE FROM institution_main_research_directions
WHERE id = $1;

-- name: GetInstitutionMainResearchDirectionByID :one
SELECT *
FROM institution_main_research_directions
WHERE id = $1;

-- name: GetInstitutionMainResearchDirectionsByInstitutionIDAndLanguage :many
SELECT *
FROM institution_main_research_directions
WHERE institution_id = $1 AND language_code = $2;

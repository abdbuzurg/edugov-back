-- name: CreateInstitutionResearchSupportInfrastructure :one
INSERT INTO institution_research_support_infrastructures(
  institution_id,
  language_code,
  research_support_infrastructure_title,
  research_support_infrastructure_type,
  tin_of_legal_entity
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionResearchSupportInfrastructure :one
UPDATE institution_research_support_infrastructures
SET
  research_support_infrastructure_title = COALESCE($1, research_support_infrastructure_title),
  research_support_infrastructure_type = COALESCE($2, research_support_infrastructure_type),
  tin_of_legal_entity = COALESCE($3, tin_of_legal_entity),
  updated_at = now()
WHERE id = $4
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionResearchSupportInfrastructure :exec
DELETE FROM institution_research_support_infrastructures
WHERE id = $1;

-- name: GetInstitutionResearchSupportInfrastructureByID :one
SELECT *
FROM institution_research_support_infrastructures
WHERE id = $1;

-- name: GetInstitutionResearchSupportInfrastructuresByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_research_support_infrastructures
WHERE institution_id = $1 AND language_code = $2;

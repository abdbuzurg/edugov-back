-- name: CreateEmployeeParticipationInProfessionalCommunity :one
INSERT INTO employee_participation_in_professional_communities(
  employee_id,
  language_code,
  professional_community_title,
  role_in_professional_community
) VALUES (
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: UpdateEmployeeParticipationInProfessionalCommunity :one
UPDATE employee_participation_in_professional_communities 
SET 
  professional_community_title = COALESCE($1, professional_community_title),
  role_in_professional_community = COALESCE($2, role_in_professional_community),
  updated_at = now()
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteEmployeeParticipationInProfessionalCommunity :exec
DELETE FROM employee_participation_in_professional_communities
WHERE id = $1;

-- name: GetEmployeeParticipationInProfessionalCommunityByID :one
SELECT *
FROM employee_participation_in_professional_communities
WHERE id = $1;

-- name: GetEmployeeParticipationInProfessionalCommunitysByEmployeeIDAndLanguageCode :many
SELECT *
FROM employee_participation_in_professional_communities
WHERE employee_id = $1 AND language_code = $2;

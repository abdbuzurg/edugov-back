-- name: CreateInstitutionAchievement :one
INSERT INTO institution_achievements(
  institution_id,
  language_code,
  achievement_title,
  achievement_type,
  date_recieved,
  given_by,
  link_to_file,
  description
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8  
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionAchievement :one
UPDATE institution_achievements
SET
  achievement_title = COALESCE($1, achievement_title),
  achievement_type = COALESCE($2, achievement_type),
  date_recieved = COALESCE($3, date_recieved),
  given_by = COALESCE($4, given_by),
  link_to_file = COALESCE($5, link_to_file),
  description = COALESCE($6, description),
  updated_at = now()
WHERE id = $7
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionAchivement :exec
DELETE FROM institution_achievements
WHERE id = $1;

-- name: GetInstitutionAchievementByID :one
SELECT *
FROM institution_achievements
WHERE id = $1;

-- name: GetInstitutionAchievementsByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_achievements
WHERE institution_id = $1 AND language_code = $2;

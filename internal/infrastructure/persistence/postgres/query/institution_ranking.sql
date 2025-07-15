-- name: CreateInstitutionRanking :one
INSERT INTO institution_rankings(
  institution_id,
  language_code,
  ranking_title,
  ranking_type,
  date_recieved,
  ranking_agency,
  link_to_ranking_file,
  description,
  link_to_ranking_agency
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, created_at, updated_at;

-- name: UpdateInstitutionRanking :one
UPDATE institution_rankings
SET
  ranking_title = COALESCE($1, ranking_title),
  ranking_type = COALESCE($2, ranking_type),
  date_recieved = COALESCE($3, date_recieved),
  ranking_agency = COALESCE($4, ranking_agency),
  link_to_ranking_file = COALESCE($5, link_to_ranking_file),
  description = COALESCE($6, description),
  link_to_ranking_agency = COALESCE($7, link_to_ranking_agency),
  updated_at = now()
WHERE id = $8
RETURNING id, created_at, updated_at;

-- name: DeleteInstitutionRanking :exec
DELETE FROM institution_rankings
WHERE id = $1;

-- name: GetInstitutionRankingByID :one
SELECT *
FROM institution_rankings
WHERE id = $1;

-- name: GetInstitutionRankingsByInstitutionIDAndLanguageCode :many
SELECT *
FROM institution_rankings
WHERE institution_id = $1 AND language_code = $2;

-- name: CreateUserSession :one
INSERT INTO user_session(
  user_id,
  refresh_token,
  expires_at
) VALUES (
  $1, $2, $3
) RETURNING id, created_at, updated_at;

-- name: UpdateUserSession :one
UPDATE user_session
SET 
  refresh_token = COALESCE($1, refresh_token),
  expires_at = COALESCE($2, refresh_token)
WHERE id = $3
RETURNING id, created_at, updated_at;

-- name: DeleteUserSessionByID :exec
DELETE FROM user_session
WHERE id = $1;

-- name: DeleteUserSessionByUserID :exec
DELETE FROM user_session
WHERE user_id = $1;

-- name: GetUserSessionByToken :one
SELECT *
FROM user_session
WHERE refresh_token = $1;

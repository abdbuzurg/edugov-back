-- name: CreateUser :one
INSERT INTO users(
  email,
  password_hash
) VALUES (
  $1, $2
) RETURNING id, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

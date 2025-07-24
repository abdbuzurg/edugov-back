-- name: CreateEmployee :one
INSERT INTO employees (
  unique_id
) VALUES(
  $1  
) RETURNING id, created_at, updated_at;

-- name: DeleteEmployee :exec
DELETE FROM employees 
WHERE id = $1;

-- name: GetEmployeeByID :one
SELECT *
FROM employees
WHERE id = $1;

-- name: GetEmployeeByUniqueIdentifier :one
SELECT *
FROM employees
WHERE unique_id = $1;

-- name: UpdateProfilePicture :exec
UPDATE employees
SET 
  profile_picture = COALESCE($1, profile_picture)
WHERE unique_id = $2;

-- name: GetProfilePicutreFileNameByUniqueID :one
SELECT profile_picture
FROM employees
WHERE unique_id = $1;

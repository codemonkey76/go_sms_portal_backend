-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, email_verified_at, created_at, updated_at;

-- name: GetUserById :one
SELECT id, name, email, email_verified_at, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, name, email, password, email_verified_at, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT id, name, email, email_verified_at, created_at, updated_at
FROM users
WHERE (sqlc.narg('search')::varchar IS NULL OR name ILIKE sqlc.narg('search'))
AND (sqlc.narg('search') IS NULL OR email ILIKE sqlc.narg('search'))
ORDER BY id ASC
LIMIT sqlc.arg('limit')::int
OFFSET sqlc.arg('offset')::int;

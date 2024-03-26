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
SELECT id, name, email, email_verified_at, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1;

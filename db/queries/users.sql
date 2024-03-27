-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, active, email_verified_at, created_at, updated_at;

-- name: GetUserById :one
SELECT id, name, email, active, email_verified_at, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, name, email, password, active, email_verified_at, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT id, name, email, active, email_verified_at, created_at, updated_at
FROM users
WHERE (sqlc.narg('search')::varchar IS NULL OR name ILIKE sqlc.narg('search'))
AND (sqlc.narg('search') IS NULL OR email ILIKE sqlc.narg('search'))
ORDER BY id ASC
LIMIT sqlc.arg('limit')::int
OFFSET sqlc.arg('offset')::int;

-- name: ListUserPermissions :many
SELECT DISTINCT p.name
FROM users u
JOIN role_user ru on u.id = ru.user_id
JOIN permission_role pr on ru.role_id = pr.role_id
JOIN permissions p on pr.permission_id = p.id
WHERE u.id = $1;



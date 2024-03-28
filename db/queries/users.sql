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
FROM permissions p
LEFT JOIN permission_user pu ON p.id = pu.permission_id AND pu.user_id = $1
LEFT JOIN permission_role pr ON p.id = pr.permission_id
LEFT JOIN role_user ru ON pr.role_id = ru.role_id AND ru.user_id = $1
WHERE pu.user_id = $1 OR ru.user_id = $1;

-- name: AttachPermissionToUser :one
INSERT INTO permission_user (permission_id, user_id) VALUES ($1, $2) RETURNING *;

-- name: AttachRoleToUser :one
INSERT INTO role_user (role_id, user_id) VALUES ($1, $2) RETURNING *;

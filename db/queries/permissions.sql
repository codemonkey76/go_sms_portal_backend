-- name: CreatePermission :one
INSERT INTO permissions (name) VALUES ($1) RETURNING *;

-- name: GetPermissionByName :one
SELECT * FROM permissions WHERE name = $1 LIMIT 1;

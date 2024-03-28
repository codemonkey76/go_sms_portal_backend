-- name: CreateRole :one
INSERT INTO roles (name) VALUES ($1) RETURNING *;

-- name: AttachPermissionToRole :one 
INSERT INTO permission_role (permission_id, role_id) VALUES ($1, $2) RETURNING *;

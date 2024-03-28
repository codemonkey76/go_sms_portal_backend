// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const attachPermissionToUser = `-- name: AttachPermissionToUser :one
INSERT INTO permission_user (permission_id, user_id) VALUES ($1, $2) RETURNING id, permission_id, user_id, created_at, updated_at
`

type AttachPermissionToUserParams struct {
	PermissionID int64 `json:"permission_id"`
	UserID       int64 `json:"user_id"`
}

func (q *Queries) AttachPermissionToUser(ctx context.Context, arg AttachPermissionToUserParams) (PermissionUser, error) {
	row := q.queryRow(ctx, q.attachPermissionToUserStmt, attachPermissionToUser, arg.PermissionID, arg.UserID)
	var i PermissionUser
	err := row.Scan(
		&i.ID,
		&i.PermissionID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const attachRoleToUser = `-- name: AttachRoleToUser :one
INSERT INTO role_user (role_id, user_id) VALUES ($1, $2) RETURNING id, role_id, user_id, created_at, updated_at
`

type AttachRoleToUserParams struct {
	RoleID int64 `json:"role_id"`
	UserID int64 `json:"user_id"`
}

func (q *Queries) AttachRoleToUser(ctx context.Context, arg AttachRoleToUserParams) (RoleUser, error) {
	row := q.queryRow(ctx, q.attachRoleToUserStmt, attachRoleToUser, arg.RoleID, arg.UserID)
	var i RoleUser
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, active, email_verified_at, created_at, updated_at
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRow struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Name, arg.Email, arg.Password)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Active,
		&i.EmailVerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, active, email_verified_at, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1
`

type GetUserByEmailRow struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Password        string       `json:"password"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.EmailVerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, active, email_verified_at, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1
`

type GetUserByIdRow struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

func (q *Queries) GetUserById(ctx context.Context, id int64) (GetUserByIdRow, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Active,
		&i.EmailVerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUserPermissions = `-- name: ListUserPermissions :many
SELECT DISTINCT p.name
FROM permissions p
LEFT JOIN permission_user pu ON p.id = pu.permission_id AND pu.user_id = $1
LEFT JOIN permission_role pr ON p.id = pr.permission_id
LEFT JOIN role_user ru ON pr.role_id = ru.role_id AND ru.user_id = $1
WHERE pu.user_id = $1 OR ru.user_id = $1
`

func (q *Queries) ListUserPermissions(ctx context.Context, userID int64) ([]string, error) {
	rows, err := q.query(ctx, q.listUserPermissionsStmt, listUserPermissions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
WITH filtered_users AS (
SELECT id, name, email, active, email_verified_at, created_at, updated_at
FROM users
WHERE ($3::varchar IS NULL OR name ILIKE $3)
AND ($3 IS NULL OR email ILIKE $3)
    ),
count_cte AS (
SELECT COUNT(*) as total_count
    FROM filtered_users
)
SELECT f.id, f.name, f.email, f.active, f.email_verified_at, f.created_at, f.updated_at, c.total_count
    FROM filtered_users f
    CROSS JOIN count_cte c
ORDER BY f.id ASC
LIMIT $2::int
OFFSET $1::int
`

type ListUsersParams struct {
	Offset int32          `json:"offset"`
	Limit  int32          `json:"limit"`
	Search sql.NullString `json:"search"`
}

type ListUsersRow struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
	TotalCount      int64        `json:"total_count"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Offset, arg.Limit, arg.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUsersRow{}
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Active,
			&i.EmailVerifiedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

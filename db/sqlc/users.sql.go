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
	ID              int32        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.Password)
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
	ID              int32        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Password        string       `json:"password"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
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
	ID              int32        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

func (q *Queries) GetUserById(ctx context.Context, id int32) (GetUserByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
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
FROM users u
JOIN role_user ru on u.id = ru.user_id
JOIN permission_role pr on ru.role_id = pr.role_id
JOIN permissions p on pr.permission_id = p.id
WHERE u.id = $1
`

func (q *Queries) ListUserPermissions(ctx context.Context, id int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listUserPermissions, id)
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
SELECT id, name, email, active, email_verified_at, created_at, updated_at
FROM users
WHERE ($1::varchar IS NULL OR name ILIKE $1)
AND ($1 IS NULL OR email ILIKE $1)
ORDER BY id ASC
LIMIT $3::int
OFFSET $2::int
`

type ListUsersParams struct {
	Search sql.NullString `json:"search"`
	Offset int32          `json:"offset"`
	Limit  int32          `json:"limit"`
}

type ListUsersRow struct {
	ID              int32        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Active          bool         `json:"active"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Search, arg.Offset, arg.Limit)
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

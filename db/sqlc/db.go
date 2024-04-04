// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.attachPermissionToRoleStmt, err = db.PrepareContext(ctx, attachPermissionToRole); err != nil {
		return nil, fmt.Errorf("error preparing query AttachPermissionToRole: %w", err)
	}
	if q.attachPermissionToUserStmt, err = db.PrepareContext(ctx, attachPermissionToUser); err != nil {
		return nil, fmt.Errorf("error preparing query AttachPermissionToUser: %w", err)
	}
	if q.attachRoleToUserStmt, err = db.PrepareContext(ctx, attachRoleToUser); err != nil {
		return nil, fmt.Errorf("error preparing query AttachRoleToUser: %w", err)
	}
	if q.createPermissionStmt, err = db.PrepareContext(ctx, createPermission); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePermission: %w", err)
	}
	if q.createRoleStmt, err = db.PrepareContext(ctx, createRole); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRole: %w", err)
	}
	if q.createSessionStmt, err = db.PrepareContext(ctx, createSession); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSession: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteExpiredSessionsStmt, err = db.PrepareContext(ctx, deleteExpiredSessions); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExpiredSessions: %w", err)
	}
	if q.deleteSessionByUserIdStmt, err = db.PrepareContext(ctx, deleteSessionByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSessionByUserId: %w", err)
	}
	if q.getPermissionByNameStmt, err = db.PrepareContext(ctx, getPermissionByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetPermissionByName: %w", err)
	}
	if q.getSessionByTokenStmt, err = db.PrepareContext(ctx, getSessionByToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetSessionByToken: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.listUserPermissionsStmt, err = db.PrepareContext(ctx, listUserPermissions); err != nil {
		return nil, fmt.Errorf("error preparing query ListUserPermissions: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.updateUserAvatarStmt, err = db.PrepareContext(ctx, updateUserAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserAvatar: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.attachPermissionToRoleStmt != nil {
		if cerr := q.attachPermissionToRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing attachPermissionToRoleStmt: %w", cerr)
		}
	}
	if q.attachPermissionToUserStmt != nil {
		if cerr := q.attachPermissionToUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing attachPermissionToUserStmt: %w", cerr)
		}
	}
	if q.attachRoleToUserStmt != nil {
		if cerr := q.attachRoleToUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing attachRoleToUserStmt: %w", cerr)
		}
	}
	if q.createPermissionStmt != nil {
		if cerr := q.createPermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPermissionStmt: %w", cerr)
		}
	}
	if q.createRoleStmt != nil {
		if cerr := q.createRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRoleStmt: %w", cerr)
		}
	}
	if q.createSessionStmt != nil {
		if cerr := q.createSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSessionStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteExpiredSessionsStmt != nil {
		if cerr := q.deleteExpiredSessionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExpiredSessionsStmt: %w", cerr)
		}
	}
	if q.deleteSessionByUserIdStmt != nil {
		if cerr := q.deleteSessionByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSessionByUserIdStmt: %w", cerr)
		}
	}
	if q.getPermissionByNameStmt != nil {
		if cerr := q.getPermissionByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPermissionByNameStmt: %w", cerr)
		}
	}
	if q.getSessionByTokenStmt != nil {
		if cerr := q.getSessionByTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSessionByTokenStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.listUserPermissionsStmt != nil {
		if cerr := q.listUserPermissionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUserPermissionsStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.updateUserAvatarStmt != nil {
		if cerr := q.updateUserAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserAvatarStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                         DBTX
	tx                         *sql.Tx
	attachPermissionToRoleStmt *sql.Stmt
	attachPermissionToUserStmt *sql.Stmt
	attachRoleToUserStmt       *sql.Stmt
	createPermissionStmt       *sql.Stmt
	createRoleStmt             *sql.Stmt
	createSessionStmt          *sql.Stmt
	createUserStmt             *sql.Stmt
	deleteExpiredSessionsStmt  *sql.Stmt
	deleteSessionByUserIdStmt  *sql.Stmt
	getPermissionByNameStmt    *sql.Stmt
	getSessionByTokenStmt      *sql.Stmt
	getUserByEmailStmt         *sql.Stmt
	getUserByIdStmt            *sql.Stmt
	listUserPermissionsStmt    *sql.Stmt
	listUsersStmt              *sql.Stmt
	updateUserAvatarStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                         tx,
		tx:                         tx,
		attachPermissionToRoleStmt: q.attachPermissionToRoleStmt,
		attachPermissionToUserStmt: q.attachPermissionToUserStmt,
		attachRoleToUserStmt:       q.attachRoleToUserStmt,
		createPermissionStmt:       q.createPermissionStmt,
		createRoleStmt:             q.createRoleStmt,
		createSessionStmt:          q.createSessionStmt,
		createUserStmt:             q.createUserStmt,
		deleteExpiredSessionsStmt:  q.deleteExpiredSessionsStmt,
		deleteSessionByUserIdStmt:  q.deleteSessionByUserIdStmt,
		getPermissionByNameStmt:    q.getPermissionByNameStmt,
		getSessionByTokenStmt:      q.getSessionByTokenStmt,
		getUserByEmailStmt:         q.getUserByEmailStmt,
		getUserByIdStmt:            q.getUserByIdStmt,
		listUserPermissionsStmt:    q.listUserPermissionsStmt,
		listUsersStmt:              q.listUsersStmt,
		updateUserAvatarStmt:       q.updateUserAvatarStmt,
	}
}

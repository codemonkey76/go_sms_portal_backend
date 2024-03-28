// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"database/sql"
	"time"
)

type Contact struct {
	ID         int64          `json:"id"`
	Phone      string         `json:"phone"`
	FirstName  sql.NullString `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	Company    sql.NullString `json:"company"`
	CustomerID int64          `json:"customer_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
}

type ContactList struct {
	ID        int64        `json:"id"`
	ContactID int64        `json:"contact_id"`
	ListID    int64        `json:"list_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Customer struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	SenderID  string       `json:"sender_id"`
	Active    bool         `json:"active"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CustomerUser struct {
	ID         int64        `json:"id"`
	CustomerID int64        `json:"customer_id"`
	UserID     int64        `json:"user_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type List struct {
	ID         int64        `json:"id"`
	Name       string       `json:"name"`
	CustomerID int64        `json:"customer_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type Message struct {
	ID         int64          `json:"id"`
	Body       string         `json:"body"`
	Segments   int32          `json:"segments"`
	From       sql.NullString `json:"from"`
	To         sql.NullString `json:"to"`
	Status     sql.NullString `json:"status"`
	Sid        sql.NullString `json:"sid"`
	Archived   bool           `json:"archived"`
	SenderID   int64          `json:"sender_id"`
	CustomerID int64          `json:"customer_id"`
	SentAt     sql.NullTime   `json:"sent_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
}

type Permission struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type PermissionRole struct {
	ID           int64        `json:"id"`
	PermissionID int64        `json:"permission_id"`
	RoleID       int64        `json:"role_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

type PermissionUser struct {
	ID           int64        `json:"id"`
	PermissionID int64        `json:"permission_id"`
	UserID       int64        `json:"user_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

type Role struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type RoleUser struct {
	ID        int64        `json:"id"`
	RoleID    int64        `json:"role_id"`
	UserID    int64        `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Session struct {
	ID           string         `json:"id"`
	IpAddress    sql.NullString `json:"ip_address"`
	UserAgent    sql.NullString `json:"user_agent"`
	Payload      string         `json:"payload"`
	LastActivity int64          `json:"last_activity"`
	UserID       int64          `json:"user_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

type Template struct {
	ID         int64        `json:"id"`
	Name       string       `json:"name"`
	Content    string       `json:"content"`
	CustomerID int64        `json:"customer_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type User struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	Password        string       `json:"password"`
	Active          bool         `json:"active"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

package sqlc

import (
	"encoding/json"
	"time"
)

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User // Use an alias to prevent recursion with MarshalJSON

	// Convert sql.NullTime fields to a pointer of time.Time or nil if not valid
	var emailVerifiedAt *time.Time
	if u.EmailVerifiedAt.Valid {
		emailVerifiedAt = &u.EmailVerifiedAt.Time
	}

	var updatedAt *time.Time
	if u.UpdatedAt.Valid {
		updatedAt = &u.UpdatedAt.Time
	}

	return json.Marshal(&struct {
		*Alias
		EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
		UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	}{
		Alias:           (*Alias)(&u),
		EmailVerifiedAt: emailVerifiedAt,
		UpdatedAt:       updatedAt,
	})
}

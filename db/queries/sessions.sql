-- name: CreateSession :one
INSERT INTO sessions (id, ip_address, user_agent, payload, last_activity, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetSessionByToken :one
SELECT * FROM sessions WHERE id = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions WHERE last_activity > $1;

-- name: DeleteSessionByUserId :exec
DELETE FROM sessions WHERE user_id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY username;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: CreateUser :exec
INSERT INTO users (email, username, password_hash, role) VALUES (?, ?, ?, ?);
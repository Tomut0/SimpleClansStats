-- name: ListUsers :many
SELECT * FROM users ORDER BY username;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

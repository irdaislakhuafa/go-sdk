-- USERS
-- name: CreateUser :execresult
INSERT INTO users (name, email, age) VALUES (?, ?, ?);

-- name: ListUser :many
SELECT u.* FROM users AS u;

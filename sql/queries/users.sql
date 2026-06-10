-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1;

-- name: DropUserTable :exec
DROP TABLE IF EXISTS users;

-- name: CreateUserTable :exec
CREATE TABLE users (
  id TEXT UNIQUE PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT UNIQUE NOT NULL
);

-- name: GetUsers :many
SELECT name FROM users;


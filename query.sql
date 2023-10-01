CREATE TYPE public.status AS ENUM ('removed', 'activate');

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: SignInByEmail :one
SELECT * FROM users WHERE email = $1 AND "password"=$2 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, "password", username ) VALUES ( $1, $2, $3 ) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users order by created_at;


-- name: CreateUser :one
INSERT INTO users (username, passwd, email, sex, age)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, passwd, email, sex, age;
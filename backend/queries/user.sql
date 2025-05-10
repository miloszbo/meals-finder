-- name: LoginUserWithUsername :one
SELECT username, passwdhash FROM users WHERE username = $1;  

-- name: CreateUser :exec
INSERT INTO users (
    username,
    passwdhash,
    email,
    name,
    surname,
    phone_number,
    age,
    sex
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
);
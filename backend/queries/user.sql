-- name: LoginUserWithUsername :one
SELECT username, passwdhash FROM users WHERE username = $1;  

-- name: CreateUser :exec
INSERT INTO users (
    username,
    passwdhash,
    email,
    phone_number,
    age,
    sex
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: GetUserData :one
SELECT username, created_at, email, name, surname, phone_number, age, sex, weight, height, BMI FROM users WHERE users.username = $1;

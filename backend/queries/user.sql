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

-- name: GetUserTags :many
SELECT tag_id FROM users_tags WHERE username = $1;

-- name: InsertUserTag :exec
INSERT INTO users_tags (username, tag_id) VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: DeleteUserTag :exec
DELETE FROM users_tags WHERE username = $1 AND tag_id = $2;

-- name: UpdateUserSettings :exec
UPDATE users
SET
    email        = CASE WHEN sqlc.arg('email')        = ''  THEN email        ELSE sqlc.arg('email')        END,
    name         = CASE WHEN sqlc.arg('name')         = ''  THEN name         ELSE sqlc.arg('name')         END,
    surname      = CASE WHEN sqlc.arg('surname')      = ''  THEN surname      ELSE sqlc.arg('surname')      END,
    phone_number = CASE WHEN sqlc.arg('phone_number') = ''  THEN phone_number ELSE sqlc.arg('phone_number') END,
    age          = CASE WHEN sqlc.arg('age')          = -1  THEN age          ELSE sqlc.arg('age')          END,
    sex          = CASE WHEN sqlc.arg('sex')          = ''  THEN sex          ELSE sqlc.arg('sex')          END,
    weight       = CASE WHEN sqlc.arg('weight')       = -1  THEN weight       ELSE sqlc.arg('weight')       END,
    height       = CASE WHEN sqlc.arg('height')       = -1  THEN height       ELSE sqlc.arg('height')       END,
    bmi          = CASE WHEN sqlc.arg('bmi')          = -1  THEN bmi          ELSE sqlc.arg('bmi')          END
WHERE username = sqlc.arg('username');



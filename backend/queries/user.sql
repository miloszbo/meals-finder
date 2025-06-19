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
INSERT INTO users_tags (username, tag_id)
SELECT @username::text AS username, t.id AS tag_id FROM tags t
JOIN tags_types tt ON tt.id = t.type_id
WHERE t.name = @tag_name::text AND tt.name = @tag_type_name::text
ON CONFLICT (username, tag_id) DO NOTHING;

-- name: DeleteUserTag :exec
DELETE FROM users_tags WHERE username = $1 AND tag_id = $2;

-- name: UpdateUserSettings :exec
UPDATE users
SET
email = CASE WHEN sqlc.arg('email')::text = ''  THEN email        ELSE sqlc.arg('email')::text        END,
name = CASE WHEN sqlc.arg('name')::text = ''  THEN name         ELSE sqlc.arg('name')::text         END,
surname = CASE WHEN sqlc.arg('surname')::text = ''  THEN surname      ELSE sqlc.arg('surname')::text      END,
phone_number = CASE WHEN sqlc.arg('phone_number')::text = ''  THEN phone_number ELSE sqlc.arg('phone_number')::text END,
age = CASE WHEN sqlc.arg('age')::int = -1  THEN age          ELSE sqlc.arg('age')::int          END,
sex = CASE WHEN sqlc.arg('sex')::text = ''  THEN sex          ELSE sqlc.arg('sex')::text          END,
weight = CASE WHEN sqlc.arg('weight')::int = -1  THEN weight       ELSE sqlc.arg('weight')::int       END,
height = CASE WHEN sqlc.arg('height')::int = -1  THEN height       ELSE sqlc.arg('height')::int       END,
bmi = CASE WHEN sqlc.arg('bmi')::int = -1  THEN bmi          ELSE sqlc.arg('bmi')::int          END
WHERE username = sqlc.arg('username')::text;
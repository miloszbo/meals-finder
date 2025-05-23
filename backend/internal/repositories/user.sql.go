// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
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
)
`

type CreateUserParams struct {
	Username    string `json:"username"`
	Passwdhash  string `json:"passwdhash"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	Age         int32  `json:"age"`
	Sex         string `json:"sex"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Username,
		arg.Passwdhash,
		arg.Email,
		arg.Name,
		arg.Surname,
		arg.PhoneNumber,
		arg.Age,
		arg.Sex,
	)
	return err
}

const loginUserWithUsername = `-- name: LoginUserWithUsername :one
SELECT username, passwdhash FROM users WHERE username = $1
`

type LoginUserWithUsernameRow struct {
	Username   string `json:"username"`
	Passwdhash string `json:"passwdhash"`
}

func (q *Queries) LoginUserWithUsername(ctx context.Context, username string) (LoginUserWithUsernameRow, error) {
	row := q.db.QueryRow(ctx, loginUserWithUsername, username)
	var i LoginUserWithUsernameRow
	err := row.Scan(&i.Username, &i.Passwdhash)
	return i, err
}

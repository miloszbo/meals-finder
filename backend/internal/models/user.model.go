package models

import "errors"

type LoginUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (lur *LoginUserRequest) Validate() error {
	if lur.Login == "" || lur.Password == "" {
		return errors.New("bad request")
	}
	return nil
}

type CreateUserRequest struct {
	Username     string `json:"username"`
	Passwdhash   string `json:"passwd"`        // maps to passwdhash column
	Email        string `json:"email"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	PhoneNumber  string `json:"phone_number"`
	Age          int32  `json:"age"`
	Sex          string `json:"sex"`
}

func (cur *CreateUserRequest) Validate() error {
	if cur.Username == "" || cur.Passwdhash == "" || cur.Email == "" || cur.PhoneNumber == "" || cur.Sex == "" || cur.Age <= 0 {
		return errors.New("missing required user fields")
	}
	return nil
}
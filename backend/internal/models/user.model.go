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
	Username    string `json:"username"`
	Passwdhash  string `json:"passwd"` // maps to passwdhash column
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Age         int32  `json:"age"`
	Sex         string `json:"sex"`
}

func (cur *CreateUserRequest) Validate() error {
	if cur.Username == "" || cur.Passwdhash == "" || cur.Email == "" || cur.PhoneNumber == "" || cur.Sex == "" || cur.Age <= 0 {
		return errors.New("missing required user fields")
	}
	return nil
}

type UpdateUserSettingsRequest struct {
	Email       string `json:"email"`        // "" = no update
	Name        string `json:"name"`         // "" = no update
	Surname     string `json:"surname"`      // "" = no update
	PhoneNumber string `json:"phone_number"` // "" = no update
	Age         int32  `json:"age"`          // -1 = no update
	Sex         string `json:"sex"`          // "" = no update
	Weight      int32  `json:"weight"`       // -1 = no update
	Height      int32  `json:"height"`       // -1 = no update
	Bmi         int32  `json:"bmi"`          // -1 = no update
}

type UserTag struct {
	Name    string
	TagType int32
}

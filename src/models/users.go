package models

import (
	"Ingress/src/validator"
	"regexp"
	"strings"
)

// User - struct for user data exchange between client and server
type User struct {
	Email    string `json:"Email" binding:"required"`
	Username string `json:"Username" binding:"required"`
}

func (u *User) checkIfEmptyRequest() bool {
	if len(strings.TrimSpace(u.Email)) == 0 || len(strings.TrimSpace(u.Username)) == 0 {
		return true
	}

	return false
}

func (u *User) checkUsername() bool {
	r, _ := regexp.Compile(`^[a-z0-9_-]{3,16}$`)
	return !(r.MatchString(u.Username))
}

func (u *User) checkEmail() bool {
	r, _ := regexp.Compile(`^(("[\w-\s]+")|([\w-]+(?:\.[\w-]+)*)|("[\w-\s]+")([\w-]+(?:\.[\w-]+)*))(@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$)|(@\[?((25[0-5]\.|2[0-4][0-9]\.|1[0-9]{2}\.|[0-9]{1,2}\.))((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.){2}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\]?$)`)
	return !(r.MatchString(u.Email))
}

//Run - generic for returning checker on User model
func (u *User) Run() interface{} {
	userCheck := &validator.UserCheck{
		IsEmpty:     u.checkIfEmptyRequest(),
		BadUsername: u.checkUsername(),
		BadEmail:    u.checkEmail(),
	}

	return userCheck
}

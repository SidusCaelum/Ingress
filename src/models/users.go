package models

import (
	"Ingress/src/db"
	"Ingress/src/validator"
	"log"
	"regexp"
	"strings"

	"github.com/globalsign/mgo"
)

// User - struct for user data exchange between client and server
type User struct {
	Email    string      `json:"Email" binding:"required"`
	Username string      `json:"Username" binding:"required"`
	DBConn   *db.Session `json:"-"`
}

func (u *User) checkIfEmptyRequest() bool {
	if len(strings.TrimSpace(u.Email)) == 0 || len(strings.TrimSpace(u.Username)) == 0 {
		return true
	}

	return false
}

func (u *User) checkUsername() bool {
	//NOTE: need to possibly rethink this regex to make it better?
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

//AddUser - add user to the database
func (u *User) AddUser() (bool, error) {
	var err error

	c := u.DBConn.DB("ingress").C("users")
	index := mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}

	if err = c.EnsureIndex(index); err != nil {
		//TODO: Come back here and handle this nicely
		panic(err)
	}

	if err = c.Insert(u.marshalJSON()); err != nil {
		if !mgo.IsDup(err) {
			//NOTE: probably shouldn't be fatal - or am i dumb and think this closes the program?
			log.Printf("Error inserting to db: %sn", err)
			return false, err
		}

		return true, err
	}

	return false, nil
}

//HACK: Needs to be done better? mgo continues to serialize DBConn with json:"-"
func (u *User) marshalJSON() (interface{}, error) {
	var tmp struct {
		Email    string `json:"Email" binding:"required"`
		Username string `json:"Username" binding:"required"`
	}

	tmp.Email = u.Email
	tmp.Username = u.Username
	return tmp, nil
}

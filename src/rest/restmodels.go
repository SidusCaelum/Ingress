package rest

import (
	"regexp"
	"reflect"
	"strings"
)

// Checker - testing
type Checker interface {
	run() interface{}
}

//Check - run the checker
func Check(c Checker) (interface{}, bool) {
	s := c.run()
	status := true

	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)

		if valueField.Interface() == true {
			//if one element of UserCheck is false status is false return
			status = false
			break
		}
	}

	return s, status
}

// User - struct for user data exchange between client and server
type User struct {
	Email    string `json:"Email" binding:"required"`
	Username string `json:"Username" binding:"required"`
}

// UserCheck - struct for checking user submission content
type UserCheck struct {
	IsEmpty     bool `json:"IsEmpty"`
	BadUsername bool `json:"BadUsername"`
	BadEmail    bool `json:"BadEmail"`
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
func (u *User) run() interface{} {
	userCheck := &UserCheck{
		IsEmpty:     u.checkIfEmptyRequest(),
		BadUsername: u.checkUsername(),
		BadEmail:    u.checkEmail(),
	}

	return userCheck
}

// Warehouse - struct for warehouse data exchange between client and server
type Warehouse struct {
	Owner string `json:"Owner" binding:"required"`
	Name  string `json:"Name" binding:"required"`
}

// WarehouseCheck - struct for checking warehouse submission content
type WarehouseCheck struct {
	IsEmpty          bool `json:"IsEmpty"`
	BadOwner bool `json:"BadUsername"`
	BadWarehouseName bool `json:"BadWarehouseName"`
}

func (w *Warehouse) checkIfEmptyRequest() bool {
	if len(strings.TrimSpace(w.Owner)) == 0 || len(strings.TrimSpace(w.Name)) == 0 {
		return true
	}

	return false
}

func (w *Warehouse) checkOwner() bool {
	//TODO: this will different for checking the owner used right now for placeholder
	r, _ := regexp.Compile(`^[a-z0-9_-]{3,16}$`)
	return !(r.MatchString(w.Owner))
}

func (w *Warehouse) checkWarehouseName() bool {
	r, _ := regexp.Compile(`^[a-z0-9_-]{3,16}$`)
	return !(r.MatchString(w.Name))
}

func(w *Warehouse) run() interface{} {
	warehouseCheck := &WarehouseCheck{
		IsEmpty: w.checkIfEmptyRequest(),
		BadOwner: w.checkOwner(),
		BadWarehouseName: w.checkWarehouseName(),
	}

	return warehouseCheck
}

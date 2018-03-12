package rest

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// NewUser - create a new user admin user
func NewUser(c *gin.Context) {
	newUser := &User{}

	if err := c.ShouldBindWith(newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, &UserCheck{
			IsEmpty:     true,
			BadUsername: false,
			BadEmail:    false,
		})

		return
	}

	status, userCheck := checkNewUserRequest(newUser)

	if status {
		c.JSON(http.StatusCreated, &userCheck)
		return
	}

	// Send back UserCheck - will be process client side.
	c.JSON(http.StatusConflict, &userCheck)
	return
}

// NewWarehouse - initalize new warehouse
func NewWarehouse(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"base.html",
		gin.H{
			"Message": "New warehouse made",
		},
	)
}

func checkNewUserRequest(newUser *User) (bool, *UserCheck) {
	status := true
	userCheck := &UserCheck{
		IsEmpty:     checkIfEmptyRequest(newUser),
		BadUsername: checkUsername(newUser.Username),
		BadEmail:    checkEmail(newUser.Email),
	}

	val := reflect.ValueOf(userCheck).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)

		if valueField.Interface() == true {
			//if one element of UserCheck is false status is false return
			status = false
			break
		}
	}

	return status, userCheck
}

func checkIfEmptyRequest(user *User) bool {
	if len(strings.TrimSpace(user.Email)) == 0 || len(strings.TrimSpace(user.Username)) == 0 {
		return true
	}

	return false
}

func checkUsername(userName string) bool {
	r, _ := regexp.Compile(`^[a-z0-9_-]{3,16}$`)
	return !(r.MatchString(userName))
}

func checkEmail(password string) bool {
	r, _ := regexp.Compile(`^(("[\w-\s]+")|([\w-]+(?:\.[\w-]+)*)|("[\w-\s]+")([\w-]+(?:\.[\w-]+)*))(@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$)|(@\[?((25[0-5]\.|2[0-4][0-9]\.|1[0-9]{2}\.|[0-9]{1,2}\.))((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.){2}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\]?$)`)
	return !(r.MatchString(password))
}

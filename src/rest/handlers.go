package rest

import (
	"log"
	"net/http"
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// NewUser - create a new user admin user
func NewUser(c *gin.Context) {
	newUser := &User{}
	if err := c.BindWith(newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, &UserCheck{
			Empty: false,
			Username: false,
			Email: false,
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
		Empty:    checkIfEmptyRequest(newUser),
		Username: checkUsername(newUser.Username),
		Email:    checkEmail(newUser.Email),
	}

	val := reflect.ValueOf(userCheck).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)

		if valueField.Interface() == false {
			//if one element of UserCheck is false status is false return
			status = false
			break
		}
	}

	return status, userCheck
}

func checkIfEmptyRequest(user *User) bool {
	if user.Email != "" || user.Username != "" {
		return true
	}

	log.Println("Requested User is empty")
	return false
}

func checkUsername(userName string) bool {
	r, _ := regexp.Compile(`^[a-z0-9_-]{3,16}$`)
	return r.MatchString(userName)
}

func checkEmail(password string) bool {
	r, _ := regexp.Compile(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
	return r.MatchString(password)
}

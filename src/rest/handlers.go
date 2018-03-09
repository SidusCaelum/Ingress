package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"fmt"

	"github.com/gin-gonic/gin"
)

// NewUser - create a new user admin user
func NewUser(c *gin.Context) {
	newUser := new(User)
	c.BindJSON(newUser)
	fmt.Printf("%+v\n", newUser)

	status, userCheck := checkNewUserRequest(newUser)

	json, err := json.Marshal(userCheck)
	if err != nil {
		log.Println(err.Error())
	}

	if status {
		c.JSON(http.StatusCreated, json)
		return
	}

	// Send back UserCheck - will be process client side.
	c.JSON(http.StatusConflict, json)
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

	fmt.Printf("%+v\n", userCheck)

	val := reflect.ValueOf(userCheck).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		fmt.Printf("ValueField: %t", valueField.Interface())

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

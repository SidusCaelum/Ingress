package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// NewUser - create a new user admin user
func NewUser(c *gin.Context) {
	newUser := new(User)
	c.BindJSON(newUser)
	status, userCheck := checkNewUserRequest(newUser)

	if status {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User created",
		})
	} else {
		json, err := json.Marshal(userCheck)
		if err != nil {
			test := err.Error()
			log.Println(test)
		}

		// Send back UserCheck - will be process client side.
		c.JSON(http.StatusConflict, json)
	}
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
		Empty: checkIfEmptyRequest(newUser),
	}

	v := reflect.ValueOf(userCheck).Elem()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == false {
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

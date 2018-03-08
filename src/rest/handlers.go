package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"base.html",
		gin.H{
			"message": "test",
		},
	)
}

// NewUser - create a new user admin user
func NewUser(c *gin.Context) {
	newUser := new(User)
	c.BindJSON(newUser)
	log.Println(newUser)
	//if newUser.Email
	c.JSON(200, gin.H{
		"message": "User created",
	})
}

// NewWarehouse - initalize new warehouse
func NewWarehouse(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"base.html",
		gin.H{
			"message": "New warehouse made",
		},
	)
}

func checkNewUser(newUser User) bool {
	if newUser.Email != nil && newUser.Username != nil {
		return true
	}

	return false
}

package rest

import (
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

type newUser struct {
}

// NewUser - create a new user admin user
func NewUser(c *gin.Context) {

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

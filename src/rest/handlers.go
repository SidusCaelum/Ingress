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

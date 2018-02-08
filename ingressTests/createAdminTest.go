package ingressTests

import (
	"rest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestValidNewUser - check if new user can be made
func TestValidNewUser(t *testing.T) {
	r := gin.Default()
	r.POST("/NewAdminUser", rest.NewAdminUser)

	w := PerformRequest(r, "POST", "/NewAdminUser", nil)

}

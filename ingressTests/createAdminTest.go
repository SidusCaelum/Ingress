package ingressTests

import (
	"rest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestValidNewUser(t *testing.T) {
	u, _ := rest.NewAdminUser(c * gin.Context)
}

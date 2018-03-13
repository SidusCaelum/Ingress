package rest

import (
	"net/http"

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

	userCheck, status := Check(newUser)

	if _, ok := userCheck.(*UserCheck); ok {

		if status {
			c.JSON(http.StatusCreated, &userCheck)
			return
		}

		// Send back UserCheck - will process client side.
		c.JSON(http.StatusConflict, &userCheck)
		return
	}

	//TODO: Handle if userCheck is not UserCheck
	//send response to the endpoint
	c.JSON(http.StatusInternalServerError, &UserCheck{
		IsEmpty:     false,
		BadUsername: false,
		BadEmail:    false,
	})

}

// NewWarehouse - initalize new warehouse
func NewWarehouse(c *gin.Context) {
	newWarehouse := &Warehouse{}

	if err := c.ShouldBindWith(newWarehouse, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, &WarehouseCheck{
			IsEmpty: true,
			BadOwner: false,
			BadWarehouseName: false,
		})

		return
	}

	warehouseCheck, status := Check(newWarehouse)

	if _, ok := warehouseCheck.(*WarehouseCheck); ok {
		if status {
			c.JSON(http.StatusCreated, &warehouseCheck)
			return
		}

		c.JSON(http.StatusConflict, &warehouseCheck)
		return
	}
	//TODO: Handle if warehouseCheck is not WarehouseCheck
	//send response to the endpoint
}

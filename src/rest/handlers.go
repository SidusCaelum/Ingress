package rest

//TODO: Change the http.status to ones that are more compliant
//TODO: ^Should build something out or find some middleware that handles this better

import (
	"Ingress/src/db"
	"Ingress/src/models"
	"Ingress/src/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//Test - test handler
// func Test(db *db.Session) gin.HandlerFunc {
// 	fn := func(c *gin.Context) {
// 		if err := db.Ping(); err != nil {
// 			log.Println(err)
// 		}

// 		x := db.DB("ingress").C("test")
// 		if err := x.Insert(
// 			models.User{
// 				Email:    "test@test.com",
// 				Username: "test",
// 			},
// 		); err != nil {
// 			log.Println("shit didn't work")
// 		}

// 		log.Println("if nothing came before this holy shit it worked")
// 	}

// 	return gin.HandlerFunc(fn)
// }

// NewUser - create a new user admin user
func NewUser(db *db.Session) gin.HandlerFunc {
	// NOTE: if this works should context be a single reference instead
	// of creating a complete new one for each handler
	fn := func(c *gin.Context) {
		newUser := &models.User{
			DBConn: db,
		}

		//TODO: handle if you get a err from unmarshalling incorrect values 234 marshalled to bool
		//This should be done client side, but double checked here
		if err := c.ShouldBindWith(newUser, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, &validator.UserCheck{
				IsEmpty:     true,
				BadUsername: false,
				BadEmail:    false,
				IsAdmin:     false,
			})

			return
		}

		userCheck, status := validator.Validate(newUser)

		//NOTE: this should be changed? Not enough information sent back
		//to clarify the err - probably pass back the err? Or handle error interally
		//to send back specific information
		if _, ok := userCheck.(*validator.UserCheck); ok {
			if status {
				if isDup, err := newUser.AddUser(); err != nil {
					if isDup {
						//Duplicate user
						c.JSON(http.StatusConflict, &userCheck)
						return
					}

					//Some other issue - probably should never reach here
					c.JSON(http.StatusUnprocessableEntity, &userCheck)
					return
				}

				//New user created with no issues
				c.JSON(http.StatusCreated, &userCheck)
				return
			}

			// Send back UserCheck - will process client side.
			c.JSON(http.StatusConflict, &userCheck)
			return
		}

		//TODO: Handle if userCheck is not UserCheck
		//send response to the endpoint
		c.JSON(http.StatusInternalServerError, &validator.UserCheck{
			IsEmpty:     false,
			BadUsername: false,
			BadEmail:    false,
			IsAdmin:     false,
		})

		return
	}

	return gin.HandlerFunc(fn)

}

// NewWarehouse - initalize new warehouse
func NewWarehouse(db *db.Session) gin.HandlerFunc {
	// NOTE: if this works should context be a single reference instead
	// of creating a complete new one for each handler
	fn := func(c *gin.Context) {
		newWarehouse := &models.Warehouse{
			DBConn: db,
		}

		//TODO: handle if you get a err from unmarshalling incorrect values 234 marshalled to bool
		//This should be done client side, but double checked here
		if err := c.ShouldBindWith(newWarehouse, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, &validator.WarehouseCheck{
				IsEmpty:          true,
				BadOwner:         false,
				BadWarehouseName: false,
			})

			return
		}

		warehouseCheck, status := validator.Validate(newWarehouse)

		//NOTE: this should be changed? Not enough information sent back
		//to clarify the err - probably pass back the err? Or handle error interally
		//to send back specific information
		if _, ok := warehouseCheck.(*validator.WarehouseCheck); ok {
			if status {
				if err := newWarehouse.AddWarehouse(); err != nil {
					/*if isDup {
						//Duplicate warehouse
						c.JSON(http.StatusConflict, &warehouseCheck)
						return
					}*/

					//Some other issue - probably should never reach here
					c.JSON(http.StatusUnprocessableEntity, &warehouseCheck)
					return
				}

				//New warehouse created with no issues
				c.JSON(http.StatusCreated, &warehouseCheck)
				return
			}

			c.JSON(http.StatusConflict, &warehouseCheck)
			return
		}

		//TODO: Handle if warehouseCheck is not WarehouseCheck
		//send response to the endpoint
		c.JSON(http.StatusInternalServerError, &validator.WarehouseCheck{
			IsEmpty:          false,
			BadOwner:         false,
			BadWarehouseName: false,
		})

		return
	}

	return gin.HandlerFunc(fn)
}

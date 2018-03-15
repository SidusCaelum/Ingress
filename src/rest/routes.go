package rest

import (
	"Ingress/src/models"

	"github.com/gin-gonic/gin"
)

type fn func(db *models.Session) gin.HandlerFunc

//Route - struct for containing route information
type Route struct {
	Method  string
	Path    string
	Handler fn
}

//Routes - array of gin.RouteInfo
type Routes []Route

var routes = Routes{
	Route{
		Method:  "POST",
		Path:    "/NewUser",
		Handler: NewUser,
	},
	Route{
		Method:  "POST",
		Path:    "/NewWarehouse",
		Handler: NewWarehouse,
	},
	Route{
		Method:  "POST",
		Path:    "/Test",
		Handler: Test,
	},
}

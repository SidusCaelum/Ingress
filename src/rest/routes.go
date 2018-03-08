package rest

import (
	gin "github.com/gin-gonic/gin"
)

//Route - struct for containing route information
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
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
}

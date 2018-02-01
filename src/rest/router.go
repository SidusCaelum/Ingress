package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// NewRouter - create new router for server. Adds all routes
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Path, route.Handler)
		case "POST":
			router.POST(route.Path, route.Handler)
		default:
			log.Printf("%s - method not recognized", route.Path)
		}
	}

	router.Static("../asset/css", "../asset/css")
	router.Static("../asset/js", "../asset/js")

	return router
}

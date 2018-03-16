package main

import (
	"Ingress/src/app"
	"Ingress/src/models"
	"log"
)

func main() {
	log.Println("Starting up")

	config := new(models.StartupConfiguration)

	a := new(app.App)
	a.Initialize(config)
	a.Run()
}

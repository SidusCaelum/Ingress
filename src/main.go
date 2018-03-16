package main

import (
	"Ingress/src/models"
	"log"
)

func main() {
	log.Println("Starting up")

	config := new(models.StartupConfiguration)

	a := new(App)
	a.Initialize(config)
	a.Run()
}

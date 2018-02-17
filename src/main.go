package main

import (
	"Ingress/src/rest"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// StartupConfiguration - where config data is read into for startup
type StartupConfiguration struct {
	Port int
}

func main() {
	log.Println("Starting up")

	// Load in the config and pipe to struct
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	config := StartupConfiguration{}
	err := decoder.Decode(&config)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Listenting on port: %d\n", config.Port)

	r := rest.NewRouter()
	r.Run()
}

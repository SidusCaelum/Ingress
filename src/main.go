package main

import (
	DB "Ingress/src/db"
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
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("error:", err)
	}

	db, err := DB.InitDB("localhost")
	if err != nil {
		log.Printf("DB connection failed: %s", err)
	}

	defer db.Close()

	fmt.Printf("Listenting on port: %d\n", config.Port)

	//HACK: passing parameter to clarify if testing or not
	r := rest.NewRouter(false, db)
	r.Run()
}

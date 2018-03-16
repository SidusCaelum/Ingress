package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// StartupConfiguration - where config data is read into for startup
type StartupConfiguration struct {
	Port   int
	DBIp   string
	DBPort string
}

//Initalize - create startup config
func (s *StartupConfiguration) Initalize() {
	// Load in the config and pipe to struct
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	config := StartupConfiguration{}
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("error:", err)
	}
}

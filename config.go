package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var conf *Configuration

// // Loads the config
func getConfig() *Configuration {
	fmt.Println("This is a config", conf)
	if conf != nil {
		return conf
	}
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf
}

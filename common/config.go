package common

import (
	"encoding/json"
	"fmt"
	"github.com/alextanhongpin/notification-challenge/model"
	"os"
)

var conf model.Configuration

// Loads the config
func GetConfig() model.Configuration {
	fmt.Println("This is a config", conf)
	if (model.Configuration{}) != conf {
		return conf
	}
	file, _ := os.Open("./conf.json")
	decoder := json.NewDecoder(file)
	configuration := model.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error:", err)
	}
	conf = configuration
	return conf
}

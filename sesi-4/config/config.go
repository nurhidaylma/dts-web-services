package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var configEnv map[string]string

func init() {
	content, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Println("config.json file not found")
		log.Println(".. using default config")
	} else {
		err = json.Unmarshal(content, &configEnv)
		if err != nil {
			log.Println("invalid config.json file")
			log.Println(".. using default config")
		}
	}
}

func GetValue(key string) string {
	value, ok := configEnv[key] // if value is empty, check config.json file
	if !ok {
		return ""
	}

	return value
}

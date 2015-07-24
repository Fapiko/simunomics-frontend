package config
import (
	"encoding/json"
	"log"
	"os"
	"io/ioutil"
)

var configLocation = "/home/ljandrew/.config/Fapiko/simunomics-frontend"
const configFilename  = "configuration.json"

type JsonConfig struct {
	data *ConfigData
}

func (config *JsonConfig) Save() {
	// Jsonify data
	jsonData, err := json.Marshal(config.data)
	if err != nil {
		log.Print("Could not marshall json")
		log.Print(err)
	}

	// Ensure directory structure exists for save location
	// TODO: Prefix home directory
	if _, err := os.Stat(configLocation); os.IsNotExist(err) {
		err := os.MkdirAll(configLocation, os.FileMode(int(0755)))
		if err != nil {
			log.Print(err)
		}
	}

	// Persist to disk
	ioutil.WriteFile(configLocation + "/" + configFilename, jsonData, 0644)
}

func (config *JsonConfig) Put(data *ConfigData) {
	config.data = data
}

func (config *JsonConfig) Get() *ConfigData {
	if config.data == nil {
		file, err := ioutil.ReadFile(configLocation + "/" + configFilename)
		if err != nil {
			log.Print(err)
		}

		err = json.Unmarshal(file, &config.data)
		logErr(err)

	}

	return config.data
}

func logErr(err error) {
	if err != nil {
		log.Print(err)
	}
}
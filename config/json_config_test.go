package config
import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"github.com/stretchr/testify/assert"
)

var dest *testing.T

func TestPut(t *testing.T) {
	dest = t

	data := &ConfigData{
		Username: "ArnoldFacepalmer",
	}

	config := &JsonConfig{}

	config.Put(data)

	if (config.data != data) {
		t.Error("Did not insert data into config")
	}
}

func TestGet(t *testing.T) {
	dest = t

	data := &ConfigData{
		Username: "ArnoldFacepalmer",
	}

	config := &JsonConfig{
		data: data,
	}

	assertEqual(data, config.Get())
}

func TestSave(t *testing.T) {
	dest = t

	setTempConfigLocation()
	defer cleanupTempConfigLocation()
	config := &JsonConfig{
		data: &ConfigData{
			Username: "ArnoldFacepalmer",
		},
	}

	config.Save()

	file, err := ioutil.ReadFile(configLocation + "/" + configFilename)
	if err != nil {
		t.Error("Could not read the config file")
		t.Error(err)
	}

	expected, err := json.Marshal(config.data)
	if err != nil {
		t.Error(err)
	}

	assertEqual(string(expected), string(file))
}

func TestGetFromFilesystem(t *testing.T) {
	setTempConfigLocation()
	defer cleanupTempConfigLocation()

	config := &JsonConfig{
		data: &ConfigData{
			Username: "ArnoldFacepalmer",
			Session: "12345",
		},
	}
	config.Save()

	newConfig := &JsonConfig{}

	assert.Equal(t, config.data, newConfig.Get())
}

func setTempConfigLocation() {
	tmpConfig, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		log.Panic(err)
	}
	configLocation = tmpConfig
}

func cleanupTempConfigLocation() {
	os.RemoveAll(configLocation)
}

func assertEqual(expected interface{}, actual interface{}, message ...string) {
	t := dest

	if expected == nil || actual == nil {
		log.Print("expected == actual: ", expected == actual)
	}

	if expected != actual {
		if len(message) != 0 {
			t.Error("Expected: ", expected, " Received: ", actual, message)
		} else {
			t.Error("Expected: ", expected, " Received: ", actual)
		}
	}
}

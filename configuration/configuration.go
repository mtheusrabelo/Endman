package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Application struct {
	Port    string            `json:"port"`
	Message map[string]string `json:"message"`
}

func init() {
	data, err := ioutil.ReadFile("configuration.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(data, &Application)
}

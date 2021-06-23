package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Controller    Controller `json:"controller"`
	PublicKeyPath string     `json:"public_key_path"`
}

type Controller struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
	Auth Auth   `json:"auth"`
}

type Auth struct {
	Token1 string `json:"token1"`
	Token2 string `json:"token2"`
	Token3 string `json:"token3"`
}

var Conf Config

func GetConfig(inputConfPath string) error {
	configPath := "./data.json"
	if inputConfPath != "" {
		configPath = inputConfPath
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	var data Config
	json.Unmarshal(file, &data)
	Conf = data
	return nil
}

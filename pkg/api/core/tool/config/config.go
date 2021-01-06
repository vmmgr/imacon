package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Controller Controller `json:"controller"`
	ImaCon     ImaCon     `json:"imacon"`
	DB         DB         `json:"db"`
	Storage    []Storage  `json:"storage"`
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

type ImaCon struct {
	Url  string `json:"url"`
	Port int    `json:"port"`
}

type AdminAuth struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type DB struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbName"`
}

type Storage struct {
	Type uint   `json:"type"`
	Path string `json:"path"`
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

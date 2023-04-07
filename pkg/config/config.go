package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//{
//"app": {
//"name": "qa-vcs-service-dtnp",
//"env": "local"
//},
//"api": {
//"host": "0.0.0.0",
//"port": "7800",
//"defaultVersion": "v1"
//},
//"log": {
//"debug": true,
//"json": false
//},
//"batch": {
//},
//"db": {
//},
//"pubsub": {
//}
//}

var ConfigCache Config

type App struct {
	Name        string `json:"name"`
	Environment string `json:"env"`
}
type Api struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	DefaultVersion string `json:"defaultVersion"`
}
type Log struct {
	Debug bool   `json:"debug"`
	JSON  bool `json:"json"`
}
type Allure struct {
	Enabled bool `json:"enabled"`
}
type Config struct {
	App    App
	Api    Api
	Log    Log
	Allure Allure
}

func GetConfig(file string) (Config, error) {
	if file == "" {
		file = "../../qa-vcs-service-dtnp.config.json"
	}

	configFile, err := os.Open(file)
	if err != nil {
		return Config{}, fmt.Errorf("opening config: %w", err)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&ConfigCache); err != nil {
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}

	return ConfigCache, nil
}



package config

import (
	"encoding/json"
	"fmt"
	"os"
)

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

// GetConfig opens and checks the file for each test
//   due to different threads, this cannot use naive/simple caching
//   This will invoke a new read at every "GetConfig" call
func GetConfig(file string) (Config, error) {
	if file == "" {
		file = "../../qa-vcs-service-dtnp.config.json"
	}

	configFile, err := os.Open(file)
	if err != nil {
		return Config{}, fmt.Errorf("opening config: %w", err)
	}

	jsonParser := json.NewDecoder(configFile)
	var c Config
	if err = jsonParser.Decode(&c); err != nil {
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}

	return c, nil
}

//
//var CacheConfig map[string]Config
//func GetConfig(file string) (Config, error) {
//	if file == "" {
//		file = "../../qa-vcs-service-dtnp.config.json"
//	}
//
//	// Does exist in cache:
//	if _, ok := CacheConfig[file]; !ok {
//		configFile, err := os.Open(file)
//		if err != nil {
//			return Config{}, fmt.Errorf("opening config: %w", err)
//		}
//
//		jsonParser := json.NewDecoder(configFile)
//		var c Config
//		if err = jsonParser.Decode(&c); err != nil {
//			return Config{}, fmt.Errorf("parsing config: %w", err)
//		}
//
//		CacheConfig[file] = c
//		fmt.Printf("\n\nGet Config: %s -- New Cache\n\n", file)
//	} else {
//		fmt.Printf("\n\nGet Config: %s -- CACHE HIT\n\n", file)
//	}
//
//	//pp, _ := json.MarshalIndent(ConfigCache, "", " ")
//	//fmt.Printf("Pretty Print:\n%s\n", string(pp))
//
//	return CacheConfig[file], nil
//}
//
//

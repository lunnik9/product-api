package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configs struct {
	Postgres PostgresConfig `json:"postgres"`
	Http     HttpConfig     `json:"http"`
}

type PostgresConfig struct {
	Url string `json:"url"`
}

type HttpConfig struct {
	Port string `json:"port"`
}

var AllConfigs *Configs

func GetConfigs() error {

	var filePath string
	if os.Getenv("config") == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filePath = pwd + "/config/config.json"
	} else {
		filePath = os.Getenv("config")
	}
	file, err := os.Open(filePath)

	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	var configs Configs
	err = decoder.Decode(&configs)

	if err != nil {
		return err
	}
	AllConfigs = &configs
	return nil
}

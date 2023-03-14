package configuration

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	BaseUrl   string
	ApiToken  string
	VerifySSL bool
}

func Read() *Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Cannot open configuration file: ", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Cannot load configuration file: ", err)
	}

	cfg := &Config{}
	err = json.Unmarshal(bytes, cfg)
	if err != nil {
		log.Fatal("Incorrect configuration format: ", err)
	}
	return cfg
}

package config

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type ServerConfiguration struct {
	Port int `yaml:"port"`
}

type Configuration struct {
	Server ServerConfiguration `yaml:"server"`
}

const defaultPort int = 3000
const defaultConfigFileLocation = "./config.yaml"

func LoadConfiguration() (Configuration, error) {
	var configFileLocation string = os.Getenv("CODE_PULSE_CONFIG_FILE")
	if configFileLocation == "" {
		log.Printf("No configuration file specified using default: %s", defaultConfigFileLocation)
		configFileLocation = defaultConfigFileLocation
	}

	fileContent, err := os.ReadFile(configFileLocation)
	log.Printf("Reading config file at: %s\n", configFileLocation)
	if err != nil {
		log.Printf("Unable to read configuration file: %s\n", configFileLocation)
		return Configuration{}, err
	}
	log.Printf("Configuration file read from: %s\n", configFileLocation)

	conf := Configuration{
		Server: ServerConfiguration{
			Port: defaultPort,
		},
	}
	err = yaml.Unmarshal(fileContent, &conf)

	if err != nil {
		log.Printf("Unable to parse configuration file %s\n", configFileLocation)
		return Configuration{}, err
	}

	return conf, nil
}

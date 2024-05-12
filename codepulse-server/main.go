package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-yaml/yaml"
)

var configFileLocation string = os.Getenv("CODE_PULSE_CONFIG_FILE")

const defaultPort int = 3000

type configuration struct {
	Port int `yaml:"port"`
}

func main() {
	if configFileLocation == "" {
		log.Println("No configuration file specified, using default: ./config.yaml")
		configFileLocation = "./config.yaml"
	}
	fileContent, err := os.ReadFile(configFileLocation)
	if err != nil {
		log.Panicf("Unable to read configuration file: %s\n", configFileLocation)
	}

	conf := configuration{Port: defaultPort}
	err = yaml.Unmarshal(fileContent, &conf)

	if err != nil {
		log.Fatalf("Unable to parse configuration file %s\n", configFileLocation)
	}

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	log.Printf("Listening on port %d", conf.Port)
	err = http.ListenAndServe(":"+strconv.Itoa(conf.Port), fs)

	if err != nil {
		log.Fatalf("Unable to start server on port %d\n", conf.Port)
	}

}

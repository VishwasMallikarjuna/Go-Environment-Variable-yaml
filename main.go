package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
	Server struct {
		Port    int `yaml:"port"`
		Timeout int `yaml:"timeout"`
	} `yaml:"server"`
	Secrets struct {
		APIKey    string `yaml:"api_key"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"secrets"`
}

func main() {
	// Read the config.yml file
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Unmarshal YAML data into Config struct
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Print the loaded configuration
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("API Key: %s\n", config.Secrets.APIKey)
}

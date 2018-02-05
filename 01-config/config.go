package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//A simple program to read from a properties file during program startup
//This program uses config.json file to store the properties in JSON format
func main() {
	config := loadConfig("config.json")
	fmt.Println("Following properties loaded from config.json file:")
	fmt.Println("User Name:", config.User)
	fmt.Println("FTP Server:", config.Server)
}

//Config holds data about configuration information for this program
type Config struct {
	User            string
	Password        string
	Server          string
	LocalPath       string
	FtpRemotePath   string
	ConnectionRetry int
}

func loadConfig(file string) Config {
	var cfg Config
	configFile, err := os.Open(file)
	if err != nil {
		log.Fatal("Could not open config file", file)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		log.Fatal("JSON unmarshal error")
	}
	return cfg
}

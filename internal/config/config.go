package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/heretic1321/gator/internal/errorhandler"
)


type Config struct {
	DBURL string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
  
	filePath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening config file")
		return Config{}, err
	}

	config, err := io.ReadAll(configFile)
	if err != nil{
		return Config{}, errorhandler.Handle(err) 
	}
	var unmarshalledConfig Config
	json.Unmarshal(config,&unmarshalledConfig)

	return unmarshalledConfig,err
}

func New() (Config,error){
	config, err := Read()
	if err != nil {
		return Config{}, err
	}

	return config, nil
}


func (c *Config) SetUser(name string) error{
	c.CurrentUsername = name
	err := c.writeToConfig()
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error){	
	homeDir, err := os.UserHomeDir()
	if err != nil{
		return "", err 
	}
	return homeDir + "/" +configFileName, nil
}


func (c *Config) writeToConfig() error{
	jsonData,err := json.Marshal(c)
	if err != nil {
		return err
	}
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

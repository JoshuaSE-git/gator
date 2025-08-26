package config

import (
	"encoding/json"
	"os"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user

	err := write(*cfg)
	if err != nil {
		return err
	}

	return nil
}

func Read() (*Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigFilePath() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := homeDirectory + "/" + configFileName

	return configPath, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, configData, 0600)
	if err != nil {
		return err
	}

	return nil
}

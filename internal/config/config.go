package config

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func ConfigPath() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("%v", err)
		return ""
	}
	path := homePath + "/.gatorconfig.json"
	return path
}

func (c Config) SetUser(user string) error {
	c.DbURL = "postgres://example"
	c.CurrentUserName = user
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Marshal Error: %v", err)
	}
	path := ConfigPath()
	if err := os.WriteFile(path, data, 0666); err != nil {
		return fmt.Errorf("WriteFile Error: %v", err)
	}
	return nil
}

func Read() (Config, error) {
	config := Config{}
	path := ConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}
	return config, nil
}

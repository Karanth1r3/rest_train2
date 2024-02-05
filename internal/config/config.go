package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	Config struct {
		HTTPServer `json:"http"`
		Logger     Logger `json:"logger"`
	}

	HTTPServer struct {
		Address string `json:"address"`
		//Port    int    `json:"port"`
	}

	Logger struct {
		LogLevel string `json:"loglevel"`
	}
)

// ReadConfig tries to read data from config file & return Config object if successfull
func ReadConfig(filePath string) (*Config, error) {
	const desc = "config.ReadConfig()"
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%s - could not read config file: %w", desc, err)
	}

	var cfg Config
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%s - could not unmarshal config: %w", desc, err)
	}
	return &cfg, nil
}

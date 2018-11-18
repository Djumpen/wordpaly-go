package cfg

import (
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	Version string `json:"version"`
}

func Load(configFileName string) (*Config, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	config := &Config{
	// Credentials:  &Credentials{},
	}
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		return nil, fmt.Errorf("cfg.go: failed to find file: err = %v", err)
	}
	configFile, err := os.Open(configFileName)
	if err != nil {
		return nil, fmt.Errorf("cfg.go: failed to open file: err = %v", err)
	}
	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("cfg.go: failed to decode file: err = %v", err)
	}

	return config, nil
}

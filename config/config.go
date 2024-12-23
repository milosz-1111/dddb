package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port int `json:"port"`

	// NoMaxCap and MaxCap are settings used
	// to determine maximum amount of elements,
	// that can be present in the database at once.
	NoMaxCap bool `json:"no_max_cap"`
	MaxCap   int  `json:"max_cap"`

	// NoMaxSize and MaxSize are needed to control
	// maximum size of values in the database.
	NoMaxSize bool `json:"no_max_size"`
	MaxSize   int  `json:"max_size"`
}

func Default() *Config {
	return &Config{
		Port:      8080,
		NoMaxCap:  false,
		MaxCap:    1024,
		NoMaxSize: false,
		MaxSize:   1024,
	}
}

func (c *Config) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("couldn't create the config file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(*c); err != nil {
		return fmt.Errorf("couldn't encode to json: %w", err)
	}

	return nil
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't open the config file: %w", err)
	}

	defer file.Close()

	var c Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("couldn't decode using json: %w", err)
	}

	return &c, nil
}

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DataBase struct {
		UserName string `yaml:"username"`
		PassWord string `yaml:"password"`
		HostName string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"database"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

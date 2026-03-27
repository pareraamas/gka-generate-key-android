package main

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Alias    string `yaml:"alias"`
	Password string `yaml:"password"`
	CN       string `yaml:"cn"`
	OU       string `yaml:"ou"`
	O        string `yaml:"o"`
	L        string `yaml:"l"`
	ST       string `yaml:"st"`
	C        string `yaml:"c"`
}

func parseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

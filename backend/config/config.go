package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Status  string `yaml:"status"`
		Shorten string `yaml:"shorten"`
	} `yaml:"server"`
	Database struct {
		Path string `yaml:"path"`
		Type string `yaml:"type"`
	} `yaml:"database"`
}

func GetConfig() Config {
	f, err := os.Open("config/config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var c Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}

	return c
}

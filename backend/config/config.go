package config

import (
	"gopkg.in/yaml.v2"
	"goshort/backend/exception"
	"os"
)

const Path = "config/config.yml"

type Config struct {
	Server struct {
		Port            string `yaml:"port"`
		RootEndpoint    string `yaml:"rootEndpoint"`
		StatusEndpoint  string `yaml:"statusEndpoint"`
		ShortenEndpoint string `yaml:"shortenEndpoint"`
	} `yaml:"server"`
	Database struct {
		Path string `yaml:"path"`
		Type string `yaml:"type"`
	} `yaml:"database"`
	ShortUrl struct {
		Length int `yaml:"length"`
	} `yaml:"shortUrl"`
}

func GetConfig() Config {
	f, err := os.Open("backend/" + Path)
	if err != nil {
		f, err = os.Open("../" + Path)
		if err != nil {
			exception.FatalInternal(err)
		}
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		exception.FatalInternal(err)
	}

	return cfg
}

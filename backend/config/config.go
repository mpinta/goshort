package config

import (
	"github.com/mpinta/goshort/backend/exception"
	"github.com/mpinta/goshort/backend/utils"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Port            string `yaml:"port"`
		Host            string `yaml:"host"`
		RootEndpoint    string `yaml:"rootEndpoint"`
		StatusEndpoint  string `yaml:"statusEndpoint"`
		ShortenEndpoint string `yaml:"shortenEndpoint"`
		LimitEndpoint   string `yaml:"limitEndpoint"`
		FindEndpoint    string `yaml:"findEndpoint"`
	} `yaml:"server"`
	Database struct {
		Type string `yaml:"type"`
	} `yaml:"database"`
	ShortUrl struct {
		Length int `yaml:"length"`
	} `yaml:"shortUrl"`
}

func GetConfig() Config {
	f, err := os.Open(utils.GetApplicationPath() + "/config.yml")
	if err != nil {
		exception.FatalInternal(err)
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

package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"database"`
	Rabbit struct {
		Dsn      string `yaml:"dsn"`
		Exchange string `yaml:"exchange"`
		Queue    string `yaml:"queue"`
		Tag      string `yaml:"tag"`
	} `yaml:"rabbit"`
}

var ErrUnsupportedType = errors.New("unsupported type")

func NewConfig(configPath string) (*Config, error) {
	if filepath.Ext(configPath) != ".yaml" {
		return nil, ErrUnsupportedType
	}

	file, _ := os.Open(configPath)
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	decoder := yaml.NewDecoder(file)

	config := &Config{}
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

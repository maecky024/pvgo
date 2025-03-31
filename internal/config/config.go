package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug  bool `yaml:"debug"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		RunLiquibase bool   `yaml:"runLiquibase"`
		Port         string `yaml:"port"`
		Host         string `yaml:"host"`
		Username     string `yaml:"user"`
		Password     string `yaml:"pass"`
		DBName       string `yaml:"dbname"`
	} `yaml:"database"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg *Config, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		processError(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

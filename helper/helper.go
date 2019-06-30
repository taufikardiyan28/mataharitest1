package helper

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type (
	Configuration struct {
		Port     int `yaml:"port"`
		Database struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DbName   string `yaml:"db_name"`
		} `yaml:"database"`
	}
)

var Config *Configuration

func ReadConfig() error {
	dat, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	Config = &Configuration{}
	err = yaml.Unmarshal(dat, Config)

	return err
}

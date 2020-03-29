package helpers

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type yamlConfig struct {
	Database struct {
		Host         string `yaml:"host"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"database"`
		Port         string `yaml:"port"`
	} `yaml:"database"`
	Jwt struct {
		DelegateSecret string `yaml:"delegateSecret"`
		ManagerSecret  string `yaml:"managerSecret"`
		AdminSecret    string `yaml:"adminSecret"`
	} `yaml:"jwt"`
}

// Config - The config variable can be used by other packages to get config data
var Config yamlConfig

// LoadConfig - Initialises the config by fetching it from the config file
func LoadConfig() error {
	filename, _ := filepath.Abs("config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config yamlConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	Config = config
	return err
}

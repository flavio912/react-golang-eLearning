package helpers

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type yamlConfig struct {
	IsTesting bool `yaml:"isTesting"`
	IsDev     bool `yaml:"isDev"`
	Database  struct {
		Host         string `yaml:"host"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"database"`
		Port         string `yaml:"port"`
	} `yaml:"database"`
	DevAdmin struct {
		Email     string `yaml:"email"`
		FirstName string `yaml:"firstName"`
		LastName  string `yaml:"lastName"`
		Password  string `yaml:"password"`
	} `yaml:"devAdmin"`
	Jwt struct {
		Secret                 string  `yaml:"secret"`
		UploadsSecret          string  `yaml:"uploadsSecret"`
		DelegateFinaliseSecret string  `yaml:"delegateFinaliseSecret"`
		CSRFSecret             string  `yaml:"csrfSecret"`
		AdminExpirationHours   float64 `yaml:"adminExpirationHours"`
	} `yaml:"jwt"`
	AWS struct {
		UploadsBucket  string `yaml:"uploadsBucket"`
		SESSendAddress string `yaml:"sesSendAddress"`
	} `yaml:"aws"`
	Imgix struct {
		BaseURL string `yaml:"baseUrl"`
	} `yaml:"imgix"`
	Sentry struct {
		DSN         string `yaml:"dsn"`
		Environment string `yaml:"environment"`
	} `yaml:"sentry"`
}

// Config - The config variable can be used by other packages to get config data
var Config yamlConfig

// LoadConfig - Initialises the config by fetching it from the config file
func LoadConfig(path string) error {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config yamlConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	if config.Jwt.Secret == config.Jwt.UploadsSecret {
		panic("Image and User token secrets should be different")
	}

	Config = config
	return err
}

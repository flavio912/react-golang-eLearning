package helpers

import (
	"os"
	"strconv"

	"github.com/golang/glog"
)

type Database struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database"`
	Port         string `yaml:"port"`
}

type DevAdmin struct {
	Email     string `yaml:"email"`
	FirstName string `yaml:"firstName"`
	LastName  string `yaml:"lastName"`
	Password  string `yaml:"password"`
}

type Jwt struct {
	Secret                 string  `yaml:"secret"`
	UploadsSecret          string  `yaml:"uploadsSecret"`
	DelegateFinaliseSecret string  `yaml:"delegateFinaliseSecret"`
	CSRFSecret             string  `yaml:"csrfSecret"`
	TokenExpirationHours   float64 `yaml:"tokenExpirationHours"`
}

type AWS struct {
	UploadsBucket  string `yaml:"uploadsBucket"`
	SESSendAddress string `yaml:"sesSendAddress"`
}

type Imgix struct {
	BaseURL string `yaml:"baseUrl"`
}

type Sentry struct {
	DSN         string `yaml:"dsn"`
	Environment string `yaml:"environment"`
}

type Stripe struct {
	PublishableKey string `yaml:"publishableKey"`
	SecretKey      string `yaml:"secretKey"`
}

type PDF struct {
	ServerURL  string `yaml:"serverUrl"`
	RequestURL string `yaml:"requestUrl"`
}

type yamlConfig struct {
	IsTesting    bool     `yaml:"isTesting"`
	IsDev        bool     `yaml:"isDev"`
	CookieDomain string   `yaml:"cookieDomain"`
	Database     Database `yaml:"database"`
	DevAdmin     DevAdmin `yaml:"devAdmin"`
	Jwt          Jwt      `yaml:"jwt"`
	AWS          AWS      `yaml:"aws"`
	Imgix        Imgix    `yaml:"imgix"`
	Sentry       Sentry   `yaml:"sentry"`
	Stripe       Stripe   `yaml:"stripe"`
	PDF          PDF      `yaml:"pdf"`
}

// Config - The config variable can be used by other packages to get config data
var Config yamlConfig

// LoadConfig - Initialises the config by fetching it from the config file
func LoadConfig() error {
	// filename, _ := filepath.Abs(path)
	// yamlFile, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }

	// var config yamlConfig
	// err = yaml.Unmarshal(yamlFile, &config)
	// if err != nil {
	// 	panic(err)
	// }

	// if !getBoolEnv("IS_TESTING") && !getBoolEnv()
	var config = yamlConfig{
		IsTesting:    getBoolEnv("IS_TESTING"),
		IsDev:        getBoolEnv("IS_DEV"),
		CookieDomain: getStringEnv("COOKIE_DOMAIN"),
		Database: Database{
			Host:         getStringEnv("DB_HOST"),
			User:         getStringEnv("DB_USER"),
			Password:     getStringEnv("DB_PASSWORD"),
			DatabaseName: getStringEnv("DB_NAME"),
			Port:         getStringEnv("DB_PORT"),
		},
		DevAdmin: DevAdmin{
			Email:     getStringEnv("DEV_ADMIN_EMAIL"),
			FirstName: getStringEnv("DEV_ADMIN_FIRST_NAME"),
			LastName:  getStringEnv("DEV_ADMIN_LAST_NAME"),
			Password:  getStringEnv("DEV_ADMIN_PASSWORD"),
		},
		Jwt: Jwt{
			Secret:                 getStringEnv("JWT_SECRET"),
			UploadsSecret:          getStringEnv("JWT_UPLOADS_SECRET"),
			DelegateFinaliseSecret: getStringEnv("JWT_FINALISE_DELEGATE_SECRET"),
			CSRFSecret:             getStringEnv("JWT_CSRF_SECRET"),
			TokenExpirationHours:   getFloatEnv("JWT_TOKEN_EXPIRATION_HOURS"),
		},
		AWS: AWS{
			UploadsBucket:  getStringEnv("AWS_UPLOADS_BUCKET"),
			SESSendAddress: getStringEnv("AWS_SES_SEND_ADDRESS"),
		},
		Imgix: Imgix{
			BaseURL: getStringEnv("IMGIX_BASE_URL"),
		},
		Sentry: Sentry{
			DSN:         getStringEnv("SENTRY_DSN"),
			Environment: getStringEnv("SENTRY_ENVIRONMENT"),
		},
		Stripe: Stripe{
			PublishableKey: getStringEnv("STRIPE_PUBLISHABLE_KEY"),
			SecretKey:      getStringEnv("STRIPE_SECRET_KEY"),
		},
		PDF: PDF{
			ServerURL:  getStringEnv("PDF_SERVER_URL"),
			RequestURL: getStringEnv("PDF_REQUEST_URL"),
		},
	}

	if config.Jwt.Secret == config.Jwt.UploadsSecret {
		panic("Image and User token secrets should be different")
	}

	Config = config
	return nil
}

func getBoolEnv(name string) bool {
	valStr := getStringEnv(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	glog.Errorf("Envvar is not bool: %s", name)
	panic("Envvar or is not bool")
}

func getStringEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	glog.Errorf("Cannot find envvar: %s", key)
	panic("Cannot find required envvar")
}

func getFloatEnv(name string) float64 {
	valueStr := getStringEnv(name)
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}

	glog.Errorf("Envvar not float: %s", name)
	panic("Envvar not float")
}

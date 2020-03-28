/*
Package database provides functions for saving and retrieving objects from the database.
*/
package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
)

var (
	gormDB *gorm.DB
)

type config struct {
	Database struct {
		Host         string `yaml:"host"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"database"`
		Port         string `yaml:"port"`
	} `yaml:"database"`
}

//SetupDatabase - Connects the database
func SetupDatabase() error {
	filename, _ := filepath.Abs("database/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	// Configure database rds or local test
	host := os.Getenv("RDS_HOSTNAME")
	port := os.Getenv("RDS_PORT")
	username := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	dbName := os.Getenv("RDS_DB_NAME")

	if host == "" {
		host = config.Database.Host
		port = config.Database.Port
		username = config.Database.User
		password = config.Database.Password
		dbName = config.Database.DatabaseName
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		dbName,
		password)
	fmt.Print(connectionString)
	db, errConnect := gorm.Open("postgres", connectionString)
	if errConnect != nil {
		return errConnect
	}

	db.AutoMigrate(&models.User{})

	gormDB = db
	return nil
}

//CreateUser - Creates a new user from required data
func CreateUser(email, password, name string) (*models.User, error) {
	if doesUserWithEmailExist(email) {
		return nil, errors.New("user with email: " + email + " already exists")
	}

	hashedPasswordBytes, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		return nil, errHash
	}

	newUser := models.User{
		Email:    email,
		Password: string(hashedPasswordBytes),
		Name:     name,
	}
	result := gormDB.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}

func doesUserWithEmailExist(email string) bool {
	user := models.User{}
	gormDB.Where("email = ?", email).First(&user)
	return user.ID != ""
}

func GetUserByCredentials(email, password string) (*models.User, error) {
	user := models.User{}
	gormDB.Where("email = ?", email).First(&user)
	if user.ID == "" {
		return nil, errors.New("no user with email " + email)
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errCompare != nil {
		return nil, errCompare
	} else {
		return &user, nil
	}
}

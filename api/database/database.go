/*
Package database provides functions for saving and retrieving objects from the database.
*/
package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	gormDB *gorm.DB
)

// Connects the database
func SetupDatabase() error {
	db, errConnect := gorm.Open("postgres",
		"host=localhost"+
			" port=5432"+
			" user=pathfinder"+
			" dbname=hackernews"+
			" password=pathfinder"+
			" sslmode=disable",
	)
	if errConnect != nil {
		return errConnect
	}

	db.AutoMigrate(&models.User{})

	gormDB = db
	return nil
}

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

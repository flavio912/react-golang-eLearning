package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
)

// User - The base model used for managers and delegates
type User struct {
	Base
	FirstName string
	LastName  string
	JobTitle  string
	Telephone string
	LastLogin time.Time
	Password  string
	Email     string `gorm:"unique"`
}

var (
	// ErrPasswordInvalid - used for ValidatePassword errors
	ErrPasswordInvalid = errors.New("Password incorrect")
)

// IUser - Interface for creating users with access tokens
type IUser interface {
	GenerateToken(string) (string, error)
	ValidatePassword(string, string) error
	FindUser(string) (IUser, error) // FindUser - Find the user by their main login method (i.e email, login_token)
	getHash() string
}

// BeforeSave - Hash the given password
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if pw, err := auth.HashPassword(user.Password); err == nil {
		scope.SetColumn("Password", pw)
	}
	return
}

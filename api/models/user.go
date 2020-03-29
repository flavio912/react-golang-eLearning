package models

import (
	"time"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"github.com/jinzhu/gorm"
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
}

// BeforeSave - Hash the given password
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if pw, err := helpers.HashPassword(user.Password); err == nil {
		scope.SetColumn("Password", pw)
	}
	return
}

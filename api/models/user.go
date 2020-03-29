package models

import (
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
}

// BeforeSave - Hash the given password
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if pw, err := auth.HashPassword(user.Password); err == nil {
		scope.SetColumn("Password", pw)
	}
	return
}

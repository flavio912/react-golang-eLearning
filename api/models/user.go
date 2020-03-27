package models

type User struct {
	Base
	Email    string
	Password string
	Name     string
}

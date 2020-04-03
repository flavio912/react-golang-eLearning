package models

// Company -
type Company struct {
	Base
	Name     string
	Managers []Manager
}

package models

import "time"

// Address table
type Address struct {
	ID           uint `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	AddressLine1 string
	AddressLine2 string
	County       string
	PostCode     string
	Country      string
}

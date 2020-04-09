package models

// CompanyRequest is a model for storing the applications to create an account
type CompanyRequest struct {
	Name      string
	Address   Address `gorm:"foreignkey:AddressID"`
	AddressID uint
	FirstName string
	LastName  string
	JobTitle  string
	Telephone string
	Email     string `gorm:"unique"`
}

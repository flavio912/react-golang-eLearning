package models

// Company -
type Company struct {
	Base
	Name      string
	Managers  []Manager
	Address   Address `gorm:"foreignkey:AddressID"`
	AddressID uint
}

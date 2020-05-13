package models

// Company -
type Company struct {
	Base
	Name      string
	Managers  []Manager  `gorm:"foreignkey:CompanyUUID"`
	Delegates []Delegate `gorm:"foreignkey:CompanyUUID"`
	Address   Address    `gorm:"foreignkey:AddressID"`
	AddressID uint
	Approved  bool
}

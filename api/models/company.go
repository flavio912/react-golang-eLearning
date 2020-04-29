package models

// Company -
type Company struct {
	Base
	Name      string
	Managers  []Manager  `gorm:"foreignkey:CompanyID"`
	Delegates []Delegate `gorm:"foreignkey:CompanyID"`
	Address   Address    `gorm:"foreignkey:AddressID"`
	AddressID uint
	Approved  bool
}

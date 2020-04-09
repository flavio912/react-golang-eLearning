package middleware

import (
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

// CreateCompanyRequest has no grant as it can be made by anyone #GRANTLESS
func CreateCompanyRequest(company gentypes.CreateCompanyInput, manager gentypes.CreateCompanyRequestManager) error {
	// Validate input
	if err := company.Validate(); err != nil {
		return err
	}
	if err := manager.Validate(); err != nil {
		return err
	}

	// Create a company request
	request := models.CompanyRequest{
		Name: company.CompanyName,
		Address: models.Address{
			AddressLine1: company.AddressLine1,
			AddressLine2: company.AddressLine2,
			County:       company.County,
			PostCode:     company.PostCode,
			Country:      company.Country,
		},
		FirstName: manager.FirstName,
		LastName:  manager.LastName,
		JobTitle:  manager.JobTitle,
		Telephone: manager.Telephone,
		Email:     manager.Email,
	}

	query := database.GormDB.Create(&request)
	if query.Error != nil {
		glog.Infof("Could not create company request: %s", query.Error.Error())
		return &errors.ErrWhileHandling
	}

	return nil
}

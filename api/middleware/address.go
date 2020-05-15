package middleware

import (
	"github.com/getsentry/sentry-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func addressToGentype(address models.Address) gentypes.Address {
	return gentypes.Address{
		ID:           address.ID,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		County:       address.County,
		PostCode:     address.PostCode,
		Country:      address.Country,
	}
}

func addressesToGentypes(addresses []models.Address) []gentypes.Address {
	var genAddresses []gentypes.Address
	for _, address := range addresses {
		genAddresses = append(genAddresses, addressToGentype(address))
	}
	return genAddresses
}

func (g *Grant) GetAddressesByIDs(ids []uint) ([]gentypes.Address, error) {
	if !g.IsAdmin {
		//TODO: Check if the users own company is in here, if so allow
		//TODO: Check if individuals address is in here, is so allow
		return []gentypes.Address{}, &errors.ErrUnauthorized
	}

	var addresses []models.Address
	query := database.GormDB.Where("id IN (?)", ids).Find(&addresses)
	if query.Error != nil {
		g.Logger.Log(sentry.LevelError, query.Error, "Unable to load addresses")
		return []gentypes.Address{}, &errors.ErrWhileHandling
	}

	return addressesToGentypes(addresses), nil
}

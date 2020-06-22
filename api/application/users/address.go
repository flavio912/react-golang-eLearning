package users

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
)

func addressesToGentypes(addresses []models.Address) []gentypes.Address {
	var genAddresses []gentypes.Address
	for _, address := range addresses {
		genAddresses = append(genAddresses, addressToGentype(address))
	}
	return genAddresses
}

func (u *usersAppImpl) GetAddressesByIDs(ids []uint) ([]gentypes.Address, error) {
	if !u.grant.IsAdmin {
		return []gentypes.Address{}, &errors.ErrUnauthorized
	}

	addresses, err := u.usersRepository.GetAddressesByIDs(ids)
	return addressesToGentypes(addresses), err
}

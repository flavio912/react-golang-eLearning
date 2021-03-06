package loader

import (
	"context"
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"github.com/graph-gophers/dataloader"
)

type companyLoader struct {
}

/*sortCompanies is a reasonably efficient way to order companies,
should be something a bit above O(2n)
*/
func sortCompanies(companies []gentypes.Company, keys dataloader.Keys) []gentypes.Company {
	var (
		k          = keys.Keys()
		companyMap = map[string]gentypes.Company{}
		sorted     = make([]gentypes.Company, len(k))
	)

	// Put companys into map of their UUIDs
	for _, company := range companies {
		companyMap[company.UUID.String()] = company
	}

	// Link keys to the companys
	for i, key := range keys {
		sorted[i] = companyMap[key.String()]
	}
	return sorted
}

func (l *companyLoader) loadCompanyBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	// Get batch from middleware
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return loadBatchError(&errors.ErrUnauthorized, n)
	}

	k := keys.Keys()
	uuidKeys := make([]gentypes.UUID, len(k))
	for i, stringK := range k {
		uuid, err := gentypes.StringToUUID(stringK)
		if err != nil {
			return loadBatchError(&errors.ErrUUIDInvalid, n)
		}
		uuidKeys[i] = uuid
	}

	userFuncs := users.NewUsersApp(grant)
	companies, err := userFuncs.GetCompaniesByUUID(uuidKeys)
	if err != nil {
		return loadBatchError(err, n)
	}

	companies = sortCompanies(companies, keys)
	res := make([]*dataloader.Result, n)
	for i, company := range companies {
		// results must be in the same order as keys
		res[i] = &dataloader.Result{Data: company}
	}
	return res
}

// LoadCompany loads Company via dataloader
func LoadCompany(ctx context.Context, uuid string) (gentypes.Company, error) {
	var company gentypes.Company
	data, err := extractAndLoad(ctx, companyLoaderKey, uuid)
	if err != nil {
		return company, err
	}
	company, ok := data.(gentypes.Company)
	if !ok {
		return company, fmt.Errorf("Wrong type: %T", data)
	}
	return company, nil
}

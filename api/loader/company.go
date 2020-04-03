package loader

import (
	"context"
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

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

func (l *companyLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	// Get batch from middleware
	grant, err := middleware.Authenticate(ctx.Value("token").(string))
	if err != nil {
		return loadBatchError(err, n)
	}

	companies, err := grant.GetCompaniesByUUID(keys.Keys())
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

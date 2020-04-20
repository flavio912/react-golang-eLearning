package loader

import (
	"context"
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"github.com/graph-gophers/dataloader"
)

type managerLoader struct {
}

/*sortManagers is a reasonably efficient way to order managers,
should be something a bit above O(2n)
*/
func sortManagers(managers []gentypes.Manager, keys dataloader.Keys) []gentypes.Manager {
	var (
		k          = keys.Keys()
		managerMap = map[string]gentypes.Manager{}
		sorted     = make([]gentypes.Manager, len(k))
	)

	// Put managers into map of their UUIDs
	for _, manager := range managers {
		managerMap[manager.UUID.String()] = manager
	}

	// Link keys to the managers
	for i, key := range keys {
		sorted[i] = managerMap[key.String()]
	}
	return sorted
}

func (l *managerLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return loadBatchError(&errors.ErrUnauthorized, n)
	}

	managers, err := grant.GetManagersByUUID(keys.Keys())
	if err != nil {
		return loadBatchError(err, n)
	}
	managers = sortManagers(managers, keys)
	res := make([]*dataloader.Result, n)
	for i, manager := range managers {
		// results must be in the same order as keys
		res[i] = &dataloader.Result{Data: manager}
	}
	return res
}

// LoadManager loads Manager via dataloader
func LoadManager(ctx context.Context, uuid string) (gentypes.Manager, error) {
	var manager gentypes.Manager
	data, err := extractAndLoad(ctx, managerLoaderKey, uuid)
	if err != nil {
		return manager, err
	}
	manager, ok := data.(gentypes.Manager)
	if !ok {
		return manager, fmt.Errorf("Wrong type: %T", data)
	}
	return manager, nil
}

type ManagerResult struct {
	Manager gentypes.Manager
	Error   error
}

func LoadManagers(ctx context.Context, uuids []string) ([]ManagerResult, error) {
	ldr, err := extract(ctx, managerLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(uuids))()

	results := make([]ManagerResult, 0, len(uuids))

	for i, d := range data {
		var e error
		if errs != nil {
			e = errs[i]
		}

		manager, ok := d.(gentypes.Manager)
		if !ok && e == nil {
			e = fmt.Errorf("Wrong type: %T", manager)
		}

		results = append(results, ManagerResult{Manager: manager, Error: e})
	}

	return results, nil
}

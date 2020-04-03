package loader

import (
	"context"
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

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

	// Get batch from middleware
	grant, err := middleware.Authenticate(ctx.Value("token").(string))
	if err != nil {
		return loadBatchError(err, n)
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

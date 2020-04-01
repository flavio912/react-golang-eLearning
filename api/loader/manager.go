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

func getManagerKeyList(loadedItems []*gentypes.Manager) []string {
	keys := make([]string, len(loadedItems))
	for i, item := range loadedItems {
		keys[i] = item.UUID
	}
	return keys
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

	res := make([]*dataloader.Result, n)
	for _, manager := range managers {
		// results must be in the same order as keys
		i := indexByString(getManagerKeyList(managers), manager.UUID)
		res[i] = &dataloader.Result{Data: manager}
	}
	return res
}

// LoadManager loads Admin via dataloader
func LoadManager(ctx context.Context, uuid string) (*gentypes.Manager, error) {
	ldr, err := extract(ctx, managerLoaderKey)
	if err != nil {
		return nil, err
	}

	v, err := ldr.Load(ctx, dataloader.StringKey(uuid))()
	if err != nil {
		return nil, err
	}
	res, ok := v.(*gentypes.Manager)
	if !ok {
		return nil, fmt.Errorf("Wrong type: %T", v)
	}

	return res, nil
}

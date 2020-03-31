package loader

import (
	"context"
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

	"github.com/graph-gophers/dataloader"
)

type adminLoader struct {
}

func getKeyList(loadedItems []*gentypes.Admin) []string {
	keys := make([]string, len(loadedItems))
	for i, item := range loadedItems {
		keys[i] = item.UUID
	}
	return keys
}

func (l *adminLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	// Get batch from middleware
	grant, err := middleware.Authenticate(ctx.Value("token").(string))
	if err != nil {
		return loadBatchError(err, n)
	}

	admins, err := grant.GetAdminsByUUID(keys.Keys())
	if err != nil {
		return loadBatchError(err, n)
	}

	res := make([]*dataloader.Result, n)
	for _, admin := range admins {
		// results must be in the same order as keys
		i := indexByString(getKeyList(admins), admin.Key())
		res[i] = &dataloader.Result{Data: admin}
	}
	return res
}

// LoadAdmin loads Admin via dataloader
func LoadAdmin(ctx context.Context, uuid string) (*gentypes.Admin, error) {
	ldr, err := extract(ctx, adminLoaderKey)
	if err != nil {
		return nil, err
	}

	v, err := ldr.Load(ctx, dataloader.StringKey(uuid))()
	if err != nil {
		return nil, err
	}
	res, ok := v.(*gentypes.Admin)
	if !ok {
		return nil, fmt.Errorf("Wrong type: %T", v)
	}

	return res, nil
}

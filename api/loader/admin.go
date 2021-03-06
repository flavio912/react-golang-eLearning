package loader

import (
	"context"
	"fmt"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"github.com/golang/glog"
	"github.com/graph-gophers/dataloader"
)

type adminLoader struct {
}

func getKeyList(loadedItems []gentypes.Admin) []gentypes.UUID {
	keys := make([]gentypes.UUID, len(loadedItems))
	for i, item := range loadedItems {
		keys[i] = item.UUID
	}
	return keys
}

func (l *adminLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)
	// Get batch from middleware
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return loadBatchError(&errors.ErrUnauthorized, n)
	}

	// Convert keys back to UUIDS
	uuids := make([]gentypes.UUID, n)
	for index, key := range keys.Keys() {
		uuids[index] = gentypes.MustParseToUUID(key)
	}

	// Get admins from application
	adminFuncs := application.NewAdminApp(grant)
	admins, err := adminFuncs.Admins(uuids)
	if err != nil {
		glog.Infof("error getting admins: %s", err.Error())
		return loadBatchError(err, n)
	}

	res := make([]*dataloader.Result, n)
	for _, admin := range admins {
		// results must be in the same order as keys
		i := indexByUUID(getKeyList(admins), admin.Key())
		res[i] = &dataloader.Result{Data: admin}
	}
	return res
}

// LoadAdmin loads Admin via dataloader
func LoadAdmin(ctx context.Context, uuid string) (gentypes.Admin, error) {
	var admin gentypes.Admin
	data, err := extractAndLoad(ctx, adminLoaderKey, uuid)
	if err != nil {
		return admin, err
	}
	admin, ok := data.(gentypes.Admin)
	if !ok {
		return admin, fmt.Errorf("Wrong type: %T", data)
	}
	return admin, nil
}

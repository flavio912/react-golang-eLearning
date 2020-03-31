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

func (l *adminLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// Get batch from DB
	jwt := ctx.Value("token").(string)

	admins, err := middleware.GetAdminsByUUID(jwt, keys.Keys())
	if err != nil {
		return []*dataloader.Result{}
	}
	res := make([]*dataloader.Result, 0)
	for _, admin := range admins {
		// results must be in the same order as keys
		res = append(res, &dataloader.Result{Data: &admin})
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
		return nil, fmt.Errorf("wrong type: %T", v)
	}
	return res, nil
}

func extract(ctx context.Context, k contextKey) (*dataloader.Loader, error) {
	res, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("cannot find a loader: %s", k)
	}
	return res, nil
}

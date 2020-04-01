package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

type contextKey string

const (
	adminLoaderKey   contextKey = "admin"
	managerLoaderKey contextKey = "manager"
)

// Init initializes and returns Map
func Init() Map {
	return Map{
		adminLoaderKey:   (&adminLoader{}).loadBatch,
		managerLoaderKey: (&managerLoader{}).loadBatch,
	}
}

// Map maps loader keys to batch-load funcs
type Map map[contextKey]dataloader.BatchFunc

// Attach attaches dataloaders to the request's context
func (m Map) Attach(ctx context.Context) context.Context {
	for k, batchFunc := range m {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFunc))
	}
	return ctx
}

func extract(ctx context.Context, k contextKey) (*dataloader.Loader, error) {
	res, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("cannot find a loader: %s", k)
	}
	return res, nil
}

func loadBatchError(err error, n int) []*dataloader.Result {
	r := &dataloader.Result{Error: err}
	res := make([]*dataloader.Result, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, r)
	}
	return res
}

func indexByString(uuids []string, uuid string) int {
	for i, v := range uuids {
		if uuid == v {
			return i
		}
	}
	panic(fmt.Sprintf("could not find %s in %v", uuid, uuids))
}

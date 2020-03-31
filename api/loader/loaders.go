package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
)

type contextKey string

const (
	adminLoaderKey contextKey = "admin"
)

// Init initializes and returns Map
func Init() Map {
	return Map{
		adminLoaderKey: (&adminLoader{}).loadBatch,
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

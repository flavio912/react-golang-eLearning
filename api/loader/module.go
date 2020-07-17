package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type moduleLoader struct {
}

func sortModules(modules []gentypes.Module, keys dataloader.Keys) []gentypes.Module {
	var (
		k         = keys.Keys()
		moduleMap = map[string]gentypes.Module{}
		sorted    = make([]gentypes.Module, len(k))
	)

	for _, module := range modules {
		moduleMap[module.UUID.String()] = module
	}

	for i, key := range keys {
		sorted[i] = moduleMap[key.String()]
	}

	return sorted
}

// loadBatch loads a batch of modules via dataloader
func (l *moduleLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	var uuids []gentypes.UUID
	for _, k := range keys.Keys() {
		uuids = append(uuids, gentypes.MustParseToUUID(k))
	}

	app := auth.AppFromContext(ctx)
	modules, err := app.CourseApp.ModulesByUUIDs(uuids)
	if err != nil {
		return loadBatchError(err, n)
	}

	modules = sortModules(modules, keys)
	res := make([]*dataloader.Result, n)
	for i, module := range modules {
		res[i] = &dataloader.Result{Data: module}
	}
	return res
}

// LoadModule loads module via dataloader
func LoadModule(ctx context.Context, uuid gentypes.UUID) (gentypes.Module, error) {
	var module gentypes.Module
	data, err := extractAndLoad(ctx, moduleLoaderKey, uuid.String())
	if err != nil {
		return module, err
	}

	module, ok := data.(gentypes.Module)
	if !ok {
		return module, fmt.Errorf("Wrong type: %T", data)
	}

	return module, nil
}

type ModuleResult struct {
	Module gentypes.Module
	Error  error
}

// LoadModules loads many modules via dataloader
func LoadModules(ctx context.Context, uuids []string) ([]ModuleResult, error) {
	ldr, err := extract(ctx, moduleLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(uuids))()

	results := make([]ModuleResult, len(uuids))

	for i, d := range data {
		var e error
		if errs != nil {
			e = errs[i]
		}

		module, ok := d.(gentypes.Module)
		if !ok && e == nil {
			e = fmt.Errorf("Wrong type: %T", module)
		}

		results[i] = ModuleResult{Module: module, Error: e}
	}

	return results, nil
}

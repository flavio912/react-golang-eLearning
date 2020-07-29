package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type testLoader struct {
}

func sortTests(tests []gentypes.Test, keys dataloader.Keys) []gentypes.Test {
	var (
		k       = keys.Keys()
		testMap = map[string]gentypes.Test{}
		sorted  = make([]gentypes.Test, len(k))
	)

	for _, test := range tests {
		testMap[test.UUID.String()] = test
	}

	for i, key := range keys {
		sorted[i] = testMap[key.String()]
	}

	return sorted
}

// loadBatch loads a batch of tests via dataloader
func (l *testLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	var uuids []gentypes.UUID
	for _, k := range keys.Keys() {
		uuids = append(uuids, gentypes.MustParseToUUID(k))
	}

	app := auth.AppFromContext(ctx)
	tests, err := app.CourseApp.TestsByUUIDs(uuids)
	if err != nil {
		return loadBatchError(err, n)
	}

	tests = sortTests(tests, keys)
	res := make([]*dataloader.Result, n)
	for i, test := range tests {
		res[i] = &dataloader.Result{Data: test}
	}
	return res
}

// LoadTest loads test via dataloader
func LoadTest(ctx context.Context, uuid gentypes.UUID) (gentypes.Test, error) {
	var test gentypes.Test
	data, err := extractAndLoad(ctx, testLoaderKey, uuid.String())
	if err != nil {
		return test, err
	}

	test, ok := data.(gentypes.Test)
	if !ok {
		return test, fmt.Errorf("Wrong type: %T", data)
	}

	return test, nil
}

type TestResult struct {
	Test  gentypes.Test
	Error error
}

// LoadTests loads many tests via dataloader
func LoadTests(ctx context.Context, uuids []string) ([]TestResult, error) {
	ldr, err := extract(ctx, testLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(uuids))()

	results := make([]TestResult, len(uuids))

	for i, d := range data {
		var e error
		if errs != nil {
			e = errs[i]
		}

		test, ok := d.(gentypes.Test)
		if !ok && e == nil {
			e = fmt.Errorf("Wrong type: %T", test)
		}

		results[i] = TestResult{Test: test, Error: e}
	}

	return results, nil
}

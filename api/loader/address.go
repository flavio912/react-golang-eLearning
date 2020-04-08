package loader

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"

	"github.com/golang/glog"
	"github.com/graph-gophers/dataloader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

type addressLoader struct {
}

/*sortAddresses is a reasonably efficient way to order Addresses,
should be something a bit above O(2n)
*/
func sortAddresses(addresses []gentypes.Address, keys dataloader.Keys) []gentypes.Address {
	var (
		k          = keys.Keys()
		addressMap = map[string]gentypes.Address{}
		sorted     = make([]gentypes.Address, len(k))
	)

	// Put addresses into map of their UUIDs
	for _, address := range addresses {
		addressMap[strconv.FormatUint(uint64(address.ID), 10)] = address
	}
	// Link keys to the addresses
	for i, key := range keys {
		sorted[i] = addressMap[key.String()]
	}
	return sorted
}

func (l *addressLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	// Get batch from middleware
	grant, err := middleware.Authenticate(ctx.Value("token").(string))
	if err != nil {
		return loadBatchError(err, n)
	}

	var ids []uint
	for _, id := range keys {
		n, err := strconv.Atoi(id.String())
		if err != nil {
			glog.Warningf("Err %s", err.Error())
			glog.Warningf("Invalid key: %s", id.String())
			return loadBatchError(&errors.ErrUnableToResolve, n)
		}
		ids = append(ids, uint(n))
	}

	addresses, err := grant.GetAddressesByIDs(ids)
	if err != nil {
		return loadBatchError(err, n)
	}
	addresses = sortAddresses(addresses, keys)
	res := make([]*dataloader.Result, n)
	for i, address := range addresses {
		// results must be in the same order as keys
		res[i] = &dataloader.Result{Data: address}
	}
	return res
}

func LoadAddress(ctx context.Context, addressID uint) (gentypes.Address, error) {
	var address gentypes.Address
	data, err := extractAndLoad(ctx, addressLoaderKey, strconv.FormatUint(uint64(addressID), 10))
	if err != nil {
		return address, err
	}
	address, ok := data.(gentypes.Address)
	if !ok {
		return address, fmt.Errorf("Wrong type: %T", data)
	}
	return address, nil
}

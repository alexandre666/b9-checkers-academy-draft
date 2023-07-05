package keeper

import (
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredenergy set a specific storedenergy in the store from its index
func (k Keeper) SetStoredenergy(ctx sdk.Context, storedenergy types.Storedenergy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredenergyKeyPrefix))
	b := k.cdc.MustMarshal(&storedenergy)
	store.Set(types.StoredenergyKey(
		storedenergy.Index,
	), b)
}

// GetStoredenergy returns a storedenergy from its index
func (k Keeper) GetStoredenergy(
	ctx sdk.Context,
	index string,

) (val types.Storedenergy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredenergyKeyPrefix))

	b := store.Get(types.StoredenergyKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredenergy removes a storedenergy from the store
func (k Keeper) RemoveStoredenergy(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredenergyKeyPrefix))
	store.Delete(types.StoredenergyKey(
		index,
	))
}

// GetAllStoredenergy returns all storedenergy
func (k Keeper) GetAllStoredenergy(ctx sdk.Context) (list []types.Storedenergy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredenergyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Storedenergy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

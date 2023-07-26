package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetStoredEnergy set a specific storedEnergy in the store from its index
func (k Keeper) SetStoredEnergy(ctx sdk.Context, storedEnergy types.StoredEnergy) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEnergyKeyPrefix))
	b := k.cdc.MustMarshal(&storedEnergy)
	store.Set(types.StoredEnergyKey(
        storedEnergy.Index,
    ), b)
}

// GetStoredEnergy returns a storedEnergy from its index
func (k Keeper) GetStoredEnergy(
    ctx sdk.Context,
    index string,
    
) (val types.StoredEnergy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEnergyKeyPrefix))

	b := store.Get(types.StoredEnergyKey(
        index,
    ))
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredEnergy removes a storedEnergy from the store
func (k Keeper) RemoveStoredEnergy(
    ctx sdk.Context,
    index string,
    
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEnergyKeyPrefix))
	store.Delete(types.StoredEnergyKey(
	    index,
    ))
}

// GetAllStoredEnergy returns all storedEnergy
func (k Keeper) GetAllStoredEnergy(ctx sdk.Context) (list []types.StoredEnergy) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredEnergyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredEnergy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

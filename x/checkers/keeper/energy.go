package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetEnergy set a specific energy in the store from its index
func (k Keeper) SetEnergy(ctx sdk.Context, energy types.Energy) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyKeyPrefix))
	b := k.cdc.MustMarshal(&energy)
	store.Set(types.EnergyKey(
        energy.Index,
    ), b)
}

// GetEnergy returns a energy from its index
func (k Keeper) GetEnergy(
    ctx sdk.Context,
    index string,
    
) (val types.Energy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyKeyPrefix))

	b := store.Get(types.EnergyKey(
        index,
    ))
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEnergy removes a energy from the store
func (k Keeper) RemoveEnergy(
    ctx sdk.Context,
    index string,
    
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyKeyPrefix))
	store.Delete(types.EnergyKey(
	    index,
    ))
}

// GetAllEnergy returns all energy
func (k Keeper) GetAllEnergy(ctx sdk.Context) (list []types.Energy) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnergyKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Energy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

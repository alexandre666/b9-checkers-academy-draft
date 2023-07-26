package keeper_test

import (
	"strconv"
	"testing"

	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/testutil/nullify"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredEnergy(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredEnergy {
	items := make([]types.StoredEnergy, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
        
		keeper.SetStoredEnergy(ctx, items[i])
	}
	return items
}

func TestStoredEnergyGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredEnergy(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredEnergy(ctx,
		    item.Index,
            
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredEnergyRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredEnergy(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredEnergy(ctx,
		    item.Index,
            
		)
		_, found := keeper.GetStoredEnergy(ctx,
		    item.Index,
            
		)
		require.False(t, found)
	}
}

func TestStoredEnergyGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredEnergy(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredEnergy(ctx)),
	)
}

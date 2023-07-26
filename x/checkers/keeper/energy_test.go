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

func createNEnergy(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Energy {
	items := make([]types.Energy, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
        
		keeper.SetEnergy(ctx, items[i])
	}
	return items
}

func TestEnergyGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNEnergy(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEnergy(ctx,
		    item.Index,
            
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEnergyRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNEnergy(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEnergy(ctx,
		    item.Index,
            
		)
		_, found := keeper.GetEnergy(ctx,
		    item.Index,
            
		)
		require.False(t, found)
	}
}

func TestEnergyGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNEnergy(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEnergy(ctx)),
	)
}

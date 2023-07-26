package keeper_test

import (
    "strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/b9lab/checkers/testutil/nullify"
	keepertest "github.com/b9lab/checkers/testutil/keeper"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEnergyQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEnergy(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetEnergyRequest
		response *types.QueryGetEnergyResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetEnergyRequest{
			    Index: msgs[0].Index,
                
			},
			response: &types.QueryGetEnergyResponse{Energy: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetEnergyRequest{
			    Index: msgs[1].Index,
                
			},
			response: &types.QueryGetEnergyResponse{Energy: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetEnergyRequest{
			    Index:strconv.Itoa(100000),
                
			},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Energy(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestEnergyQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEnergy(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEnergyRequest {
		return &types.QueryAllEnergyRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EnergyAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Energy), step)
			require.Subset(t,
            	nullify.Fill(msgs),
            	nullify.Fill(resp.Energy),
            )
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EnergyAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Energy), step)
			require.Subset(t,
            	nullify.Fill(msgs),
            	nullify.Fill(resp.Energy),
            )
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EnergyAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Energy),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EnergyAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

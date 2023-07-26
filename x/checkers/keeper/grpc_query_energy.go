package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/b9lab/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EnergyAll(c context.Context, req *types.QueryAllEnergyRequest) (*types.QueryAllEnergyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var energys []types.Energy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	energyStore := prefix.NewStore(store, types.KeyPrefix(types.EnergyKeyPrefix))

	pageRes, err := query.Paginate(energyStore, req.Pagination, func(key []byte, value []byte) error {
		var energy types.Energy
		if err := k.cdc.Unmarshal(value, &energy); err != nil {
			return err
		}

		energys = append(energys, energy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEnergyResponse{Energy: energys, Pagination: pageRes}, nil
}

func (k Keeper) Energy(c context.Context, req *types.QueryGetEnergyRequest) (*types.QueryGetEnergyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEnergy(
	    ctx,
	    req.Index,
        )
	if !found {
	    return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEnergyResponse{Energy: val}, nil
}
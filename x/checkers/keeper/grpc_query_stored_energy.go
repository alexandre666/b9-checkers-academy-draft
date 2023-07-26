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

func (k Keeper) StoredEnergyAll(c context.Context, req *types.QueryAllStoredEnergyRequest) (*types.QueryAllStoredEnergyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedEnergys []types.StoredEnergy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedEnergyStore := prefix.NewStore(store, types.KeyPrefix(types.StoredEnergyKeyPrefix))

	pageRes, err := query.Paginate(storedEnergyStore, req.Pagination, func(key []byte, value []byte) error {
		var storedEnergy types.StoredEnergy
		if err := k.cdc.Unmarshal(value, &storedEnergy); err != nil {
			return err
		}

		storedEnergys = append(storedEnergys, storedEnergy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredEnergyResponse{StoredEnergy: storedEnergys, Pagination: pageRes}, nil
}

func (k Keeper) StoredEnergy(c context.Context, req *types.QueryGetStoredEnergyRequest) (*types.QueryGetStoredEnergyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredEnergy(
	    ctx,
	    req.Index,
        )
	if !found {
	    return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredEnergyResponse{StoredEnergy: val}, nil
}
package keeper

import (
	"context"

	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredenergyAll(c context.Context, req *types.QueryAllStoredenergyRequest) (*types.QueryAllStoredenergyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedenergys []types.Storedenergy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedenergyStore := prefix.NewStore(store, types.KeyPrefix(types.StoredenergyKeyPrefix))

	pageRes, err := query.Paginate(storedenergyStore, req.Pagination, func(key []byte, value []byte) error {
		var storedenergy types.Storedenergy
		if err := k.cdc.Unmarshal(value, &storedenergy); err != nil {
			return err
		}

		storedenergys = append(storedenergys, storedenergy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredenergyResponse{Storedenergy: storedenergys, Pagination: pageRes}, nil
}

func (k Keeper) Storedenergy(c context.Context, req *types.QueryGetStoredenergyRequest) (*types.QueryGetStoredenergyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredenergy(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredenergyResponse{Storedenergy: val}, nil
}

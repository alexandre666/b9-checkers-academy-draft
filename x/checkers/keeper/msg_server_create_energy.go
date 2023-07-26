package keeper

import (
	"context"

	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEnergy(goCtx context.Context, msg *types.MsgCreateEnergy) (*types.MsgCreateEnergyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	storedEnergy := types.Energy{
		Kwh: msg.Kwh,

	}

	k.Keeper.SetEnergy(ctx, storedEnergy)

	return &types.MsgCreateEnergyResponse{}, nil
}

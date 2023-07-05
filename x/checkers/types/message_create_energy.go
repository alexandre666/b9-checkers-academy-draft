package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateEnergy = "create_energy"

var _ sdk.Msg = &MsgCreateEnergy{}

func NewMsgCreateEnergy(creator string, kwh string) *MsgCreateEnergy {
	return &MsgCreateEnergy{
		Creator: creator,
		Kwh:     kwh,
	}
}

func (msg *MsgCreateEnergy) Route() string {
	return RouterKey
}

func (msg *MsgCreateEnergy) Type() string {
	return TypeMsgCreateEnergy
}

func (msg *MsgCreateEnergy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEnergy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEnergy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

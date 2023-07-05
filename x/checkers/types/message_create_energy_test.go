package types

import (
	"testing"

	"github.com/b9lab/checkers/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateEnergy_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateEnergy
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateEnergy{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateEnergy{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

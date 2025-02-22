package types_test

import (
	"testing"

	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				SystemInfo: types.SystemInfo{
					NextId: 39,
				},
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				PlayerInfoList: []types.PlayerInfo{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				StoredenergyList: []types.Storedenergy{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				StoredEnergyList: []types.StoredEnergy{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
EnergyList: []types.Energy{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedGame",
			genState: &types.GenesisState{
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated playerInfo",
			genState: &types.GenesisState{
				PlayerInfoList: []types.PlayerInfo{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated storedenergy",
			genState: &types.GenesisState{
				StoredenergyList: []types.Storedenergy{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
	desc:     "duplicated storedEnergy",
	genState: &types.GenesisState{
		StoredEnergyList: []types.StoredEnergy{
			{
				Index: "0",
},
			{
				Index: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated energy",
	genState: &types.GenesisState{
		EnergyList: []types.Energy{
			{
				Index: "0",
},
			{
				Index: "0",
},
		},
	},
	valid:    false,
},
// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDefaultGenesisState_ExpectedInitialNextId(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			StoredGameList: []types.StoredGame{},
			SystemInfo: types.SystemInfo{
				NextId:        uint64(1),
				FifoHeadIndex: "-1",
				FifoTailIndex: "-1",
			},
			PlayerInfoList: []types.PlayerInfo{},
		},
		types.DefaultGenesis())
}

package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/b9lab/checkers/x/checkers/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group checkers queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowSystemInfo())
	cmd.AddCommand(CmdListStoredGame())
	cmd.AddCommand(CmdShowStoredGame())
	cmd.AddCommand(CmdCanPlayMove())

	cmd.AddCommand(CmdListPlayerInfo())
	cmd.AddCommand(CmdShowPlayerInfo())
	cmd.AddCommand(CmdListStoredenergy())
	cmd.AddCommand(CmdShowStoredenergy())
	cmd.AddCommand(CmdListStoredEnergy())
	cmd.AddCommand(CmdShowStoredEnergy())
cmd.AddCommand(CmdListEnergy())
	cmd.AddCommand(CmdShowEnergy())
// this line is used by starport scaffolding # 1

	return cmd
}

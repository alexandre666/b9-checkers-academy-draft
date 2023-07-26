package cli

import (
    "context"
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/b9lab/checkers/x/checkers/types"
)

func CmdListStoredEnergy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stored-energy",
		Short: "list all storedEnergy",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllStoredEnergyRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.StoredEnergyAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowStoredEnergy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stored-energy [index]",
		Short: "shows a storedEnergy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argIndex := args[0]
            
            params := &types.QueryGetStoredEnergyRequest{
                Index: argIndex,
                
            }

            res, err := queryClient.StoredEnergy(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

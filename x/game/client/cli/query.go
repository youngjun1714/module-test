package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/youngjun1714/module-test/x/game/types"
)

func NewQueryCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying command for the game module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	return cmd
}

func GetCmdQueryTotalBetting() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "totalbetting",
		Args:  cobra.NoArgs,
		Short: "Query Total Betting amount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TotalBetting(cmd.Context(), &types.QueryTotalBettingRequest{})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

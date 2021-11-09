package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/youngjun1714/module-test/x/game/types"
)

func NewTxCmd() *cobra.Command {

	gameTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "game transactions subcommand",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	gameTxCmd.AddCommand(NewBetAmountTxCmd())

	return gameTxCmd
}

func NewBetAmountTxCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "bet-amount [from_key_or_address] [Odds_or_Evens] [amount]",
		Short: "betting amount Odds and Evens",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgBetAmountRequest(clientCtx.GetFromAddress(), args[1], coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

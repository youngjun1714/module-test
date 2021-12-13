package cli

import (
	"context"

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

func NewRewardsTxCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "All-Rewards [Odds_or_Evens]",
		Short: "betting amount Rewards Odds or Evens [ default signer : game1hmnxk5d84sl33lmdyxcervdqrp9nfau0xtv9u0 ]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, types.RewardsAddress)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			var msgs []sdk.Msg

			switch args[0] {
			case "Odds":
				res, err := queryClient.OddsBetting(context.Background(), &types.QueryOddsBettingRequest{})
				if err != nil {
					return err
				}
				msgs = make([]sdk.Msg, 0, len(res.OddsInfo.Info))
				for _, info := range res.OddsInfo.Info {
					msg := types.NewMsgAllRewardsRequest(info.FromAddress, "Odds", info.Amount)

					if err := msg.ValidateBasic(); err != nil {
						return err
					}
					msgs = append(msgs, msg)
				}
			case "Evens":
				res, err := queryClient.EvensBetting(context.Background(), &types.QueryEvensBettingRequest{})

				if err != nil {
					return err
				}
				msgs = make([]sdk.Msg, 0, len(res.EvensInfo.Info))
				for _, info := range res.EvensInfo.Info {
					msg := types.NewMsgAllRewardsRequest(info.FromAddress, "Evens", info.Amount)

					if err := msg.ValidateBasic(); err != nil {
						return err
					}
					msgs = append(msgs, msg)
				}
			default:

			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

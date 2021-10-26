package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/youngjun1714/module-test/x/game/types"
)

func NewTxCmd() *cobra.Command {

	gameTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "game transactions subcommand",
		RunE:  client.ValidateCmd,
	}

	return gameTxCmd
}

func NewBetAmountTxCmd() *cobra.Command {

	return nil
}

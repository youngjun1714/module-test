package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/youngjun1714/module-test/types"
)

func NewTxCmd() *cobra.Command {

	gameTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "game transactions subcommand",
		RunE:  client.ValidateCmd,
	}

	return gameTxCmd
}

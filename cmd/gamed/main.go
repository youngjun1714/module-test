package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	game "github.com/youngjun1714/module-test/app"
	"github.com/youngjun1714/module-test/cmd/gamed/cmd"
	"github.com/youngjun1714/module-test/types"
)

func main() {
	//Set address prefix
	types.SetConfig()

	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, game.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}

}

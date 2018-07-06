package main

import (
	"fmt"
	"os"

	"github.com/prysmaticlabs/beacon-chain/node"
	"github.com/urfave/cli"
)

func startBeaconNode(ctx *cli.Context) error {
	beaconNode, err := node.New(ctx)
	if err != nil {
		return err
	}
	// starts a connection to a beacon node and kicks off every registered service.
	beaconNode.Start()
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "beacon-chain"
	app.Action = startBeaconNode
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

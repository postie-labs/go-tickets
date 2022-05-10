package cli

import (
	"fmt"

	"github.com/postie-labs/go-tickets/cli/tickets"
	"github.com/spf13/cobra"
)

const (
	DefaultLCDEndpoint            = "https://bombay-lcd.terra.dev"
	DefaultChainID                = "bombay-12"
	DefaultTimeout                = 10 // sec
	DefaultContractAddressBench32 = "terra1al87aagg7asjyceu9x8f4xj554ddzlk9q2t8ls"
)

var RootCmd = &cobra.Command{
	Use:   "tickets",
	Short: "'tickets' is a CLI tool to interact with the smart contract `cw721-tickets`",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

func init() {
	cobra.EnableCommandSorting = false
	tickets.LCDEndpoint = RootCmd.PersistentFlags().String("endpoint", DefaultLCDEndpoint, "LCDEndpoint")
	tickets.ChainID = RootCmd.PersistentFlags().String("chain-id", DefaultChainID, "Chain ID")
	tickets.Timeout = RootCmd.PersistentFlags().Int64("timeout", DefaultTimeout, "Timeout")
	RootCmd.AddCommand(
		tickets.IssueCmd,
		tickets.TransferCmd,
		tickets.ListCmd,
		tickets.GetCmd,
		tickets.QRCmd,
	)
}

package cli

import (
	"fmt"

	"github.com/postie-labs/go-tickets/cli/tickets"
	"github.com/spf13/cobra"
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
	RootCmd.AddCommand(
		tickets.IssueCmd,
		tickets.TransferCmd,
		tickets.ListCmd,
		tickets.GetCmd,
		tickets.QRCmd,
	)
}

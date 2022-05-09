package tickets

import "github.com/spf13/cobra"

var TransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfer a ticket's ownership to <recipient>",
}

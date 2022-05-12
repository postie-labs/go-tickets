package tickets

import "github.com/spf13/cobra"

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a ticket's information",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

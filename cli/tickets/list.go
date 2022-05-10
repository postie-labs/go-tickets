package tickets

import "github.com/spf13/cobra"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list issued tickets",
}

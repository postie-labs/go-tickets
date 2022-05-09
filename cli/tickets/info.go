package tickets

import "github.com/spf13/cobra"

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "get a ticket's information",
}

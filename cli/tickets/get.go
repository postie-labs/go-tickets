package tickets

import "github.com/spf13/cobra"

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a ticket's information",
}

package tickets

import "github.com/spf13/cobra"

var IssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "issue a ticket",
}

package tickets

import (
	"fmt"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var IssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "issue a ticket",
	RunE: func(cmd *cobra.Command, args []string) error {
		// prepare
		pubKey := LCDClient.PrivKey.PubKey()
		address := cosmtypes.AccAddress(pubKey.Address())
		fmt.Println(address.String())
		return nil
	},
}
